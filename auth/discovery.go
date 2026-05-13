// Copyright 2026 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const wellKnownPath = ".well-known/oauth-authorization-server"

// DiscoverAuthServer fetches the OAuth Authorization Server metadata from the
// well-known endpoint per RFC 8414. Returns a wrapper map with two keys:
// "metadata" holds the server's response as-is, and "expiry" holds an RFC 3339
// timestamp derived from HTTP cache headers to support client-side caching.
func (c *Config) DiscoverAuthServer(ctx context.Context) (map[string]any, error) {
	discoveryURL := c.AuthServerURL.ResolveReference(&url.URL{Path: wellKnownPath})

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, discoveryURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create discovery request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("discovery request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("discovery endpoint returned HTTP %d", resp.StatusCode)
	}

	var serverMetadata map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&serverMetadata); err != nil {
		return nil, fmt.Errorf("failed to parse discovery response: %w", err)
	}

	result := map[string]any{
		"metadata": serverMetadata,
	}
	if expiresAt := parseCacheExpiry(resp.Header); !expiresAt.IsZero() {
		result["expiry"] = expiresAt.Format(time.RFC3339)
	}

	return result, nil
}

// parseCacheExpiry derives an expiry time from HTTP cache headers.
// It checks Cache-Control max-age first, then falls back to the Expires header.
func parseCacheExpiry(header http.Header) time.Time {
	if cc := header.Get("Cache-Control"); cc != "" {
		for _, directive := range strings.Split(cc, ",") {
			directive = strings.TrimSpace(directive)
			if strings.HasPrefix(directive, "max-age=") {
				if seconds, err := strconv.Atoi(strings.TrimPrefix(directive, "max-age=")); err == nil {
					return time.Now().Add(time.Duration(seconds) * time.Second)
				}
			}
		}
	}

	if expires := header.Get("Expires"); expires != "" {
		if t, err := http.ParseTime(expires); err == nil {
			return t
		}
	}

	// Default to 1 week if the server sends no cache headers.
	return time.Now().Add(7 * 24 * time.Hour)
}
