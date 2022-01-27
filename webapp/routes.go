// Copyright 2017 The WPT Dashboard Project. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package webapp

import (
	"github.com/web-platform-tests/wpt.fyi/shared"
)

// RegisterRoutes adds the route handlers for the webapp.
func RegisterRoutes() {
	// GitHub OAuth login
	shared.AddRoute("/login", "login", loginHandler)
	shared.AddRoute("/logout", "logout", logoutHandler)
	shared.AddRoute("/oauth", "oauth", oauthHandler)

	// About wpt.fyi
	shared.AddRoute("/about", "about", aboutHandler)

	// Reftest analyzer
	shared.AddRoute("/analyzer", "analyzer", analyzerHandler)

	// Feature flags for wpt.fyi
	shared.AddRoute("/flags", "flags", flagsHandler)
	shared.AddRoute("/dynamic-components/wpt-env-flags.js", "flags-component", flagsComponentHandler)

	shared.AddRoute("/node_modules/{path:.*}", "components", componentsHandler)

	// Test run results, viewed by pass-rate across the browsers
	shared.AddRoute("/interop/", "interop", interopHandler)
	shared.AddRoute("/interop/{path:.*}", "interop", interopHandler)

	// A list of useful/insightful queries
	shared.AddRoute("/insights", "insights", insightsHandler)

	// List of all pending/in-flight runs
	shared.AddRoute("/status", "processor-status", processorStatusHandler)

	// List of all test runs, by SHA[0:10]
	shared.AddRoute("/runs", "test-runs", testRunsHandler)
	shared.AddRoute("/test-runs", "test-runs", testRunsHandler) // Legacy name

	// Dashboard for the compat-2021 effort.
	shared.AddRoute("/compat2021", "compat-2021", compat2021Handler)

	// Dashboard for the interop-2022 effort.
	shared.AddRoute("/interop-2022", "interop-2022", interop2022Handler)

	// Admin-only manual results upload.
	shared.AddRoute("/admin/results/upload", "admin-results-upload", adminUploadHandler)

	// Admin-only manual cache flush.
	shared.AddRoute("/admin/cache/flush", "admin-cache-flush", adminCacheFlushHandler)

	// Admin-only environment flag management
	shared.AddRoute("/admin/flags", "admin-flags", adminFlagsHandler)

	// Test run results, viewed by browser (default view)
	// For run results diff view, 'before' and 'after' params can be given.
	shared.AddRoute("/results/", "results", testResultsHandler)
	shared.AddRoute("/results/{path:.*}", "results", testResultsHandler)

	// Legacy wildcard match
	shared.AddRoute("/", "results-legacy", testResultsHandler)
	shared.AddRoute("/{path:.*}", "results-legacy", testResultsHandler)
}
