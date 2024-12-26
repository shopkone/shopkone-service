package sNav

import (
	"shopkone-service/internal/module/online/mav/mNav"
)

// 初始化默认
func (s *sNav) NavInit() (err error) {
	mainMenu := mNav.Nav{
		Title:  "Main Menu",
		Handle: "main-menu",
		Links: []mNav.NavItem{
			{
				Title: "Home",
				Url:   "/",
			},
			{
				Title: "Collections",
				Url:   "/collections",
			},
			{
				Title: "All products",
				Url:   "/collections/all",
			},
		},
	}
	if _, err = s.NavCreate(mainMenu); err != nil {
		return err
	}

	footerMenu := mNav.Nav{
		Title:  "Footer",
		Handle: "footer",
		Links: []mNav.NavItem{
			{
				Title: "About",
				Url:   "/about",
			},
			{
				Title: "Contact",
				Url:   "/contact",
			},
		},
	}
	_, err = s.NavCreate(footerMenu)
	return err
}
