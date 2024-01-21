package main

type GetModsResponse struct {
	ModName        string `json:"mod_name"`
	ModAuthor      string `json:"mod_author"`
	ModVersion     string `json:"mod_version"`
	ModDescription string `json:"mod_description"`
	ModPicture     string `json:"mod_picture"`
	ModPathName    string `json:"mod_path_name"`
}
