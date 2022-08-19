package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func svg_image(w http.ResponseWriter, r *http.Request, first_color string, second_color string, counts uint64) {
	width := (12 + (len(strconv.FormatUint(counts, 10)) * 8)) + 3

	w.Header().Add("cache-control", "max-age=0, no-cache, no-store, must-revalidate")
	w.Header().Add("content-type", "image/svg+xml")
	fmt.Fprintf(w, `
	<!--
	How-many-views-github: https://github.com/kiwimoe/hmvg
	-->
	<svg height="20" xmlns="http://www.w3.org/2000/svg">
	<rect width="95" height="20" style="fill:%s;"/>
	<text x="6" font-size="11" font-family="Verdana,monospace,sans-serif" y="15" fill="black">Profile views</text>
    <text x="5" font-size="11" font-family="Verdana,monospace,sans-serif" y="14" fill="white">Profile views</text>
    
    <g>
	<rect x="95" width="%d" height="20" style="fill:%s" rx="3" ry="3"/>
	<svg x="95" width="%d" height="20">
	<text x="51%%" font-size="11" text-anchor="middle" y="15" font-family="Verdana,monospace,sans-serif" fill="black">%d</text>
	<text x="50%%" font-size="11" text-anchor="middle" y="14" font-family="Verdana,monospace,sans-serif" fill="white">%d</text>
	</svg>
	</g>
	<rect x="93" width="3" height="20" style="fill:%s;"/>
</svg>`, first_color, width, second_color, width, counts, counts, first_color)
}
