package sNav

import (
	"shopkone-service/internal/module/online/mav/mNav"
)

// 初始化默认
func (s *sNav) NavInit() (err error) {
	mainMenu := mNav.Nav{
		Title:  "Main Menu",
		Handle: "main-menu",
	}
	mainMenuItems := []mNav.NavItem{
		{
			Title:  "Home",
			Url:    "/",
			Levels: 1,
		},
		{
			Title:  "Collections",
			Url:    "/collections",
			Levels: 1,
		},
		{
			Title:  "All products",
			Url:    "/products",
			Levels: 1,
		},
	}
	if err = s.NavCreate(mainMenu, mainMenuItems); err != nil {
		return err
	}

	footerMenu := mNav.Nav{
		Title:  "Footer",
		Handle: "footer",
	}
	footerMenuItems := []mNav.NavItem{
		{
			Title:  "Search",
			Url:    "/search",
			Levels: 1,
		},
	}
	if err = s.NavCreate(footerMenu, footerMenuItems); err != nil {
		return err
	}
	return err
}
