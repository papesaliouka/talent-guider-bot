package commands

func isAdmin(userID string) bool {
	// Implement your logic to determine if the user is an admin
	// You can use Discord's APIs or check against a list of admin user IDs

	// Sample logic to check if the user is an admin based on their ID
	adminUserIDs := []string{"787648176295378944", "adminUserID2"}
	for _, id := range adminUserIDs {
		if userID == id {
			return true
		}
	}

	return false
}
