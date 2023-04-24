/*
Copyright © 2023 Syro team <info@syro.com>
*/
package util

// Server URL
var LocalServerApiUrl = "http://localhost:1337/v1/functions"
var StagingServerApiUrl = "https://api-staging.syro.com/v1/functions"
var ProductionServerApiUrl = "https://api-production.syro.com/v1/functions"

func GetServerUrl(env string) (serverApiUrl string) {
	if env == "local" {
		return LocalServerApiUrl
	} else if env == "staging" {
		return StagingServerApiUrl
	}
	return ProductionServerApiUrl
}

// Auth API endpoints
var CliLogin = "/cli_login"
var CliValidateAccessTokenAndProjectId = "/cli_validate_access_token_and_project_id"
var CliValidateSessionToken = "/cli_validate_session_token"

// Project API endpoints
var CliValidateProjectId = "/cli_validate_project_id"
var CliProjectItems = "/cli_project_items"
var CliPullProjectItems = "/cli_pull_project_items"
