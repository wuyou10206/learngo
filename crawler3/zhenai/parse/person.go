package parse

import (
	"learngo/crawler3/engine"
	"learngo/crawler3/model"
	"regexp"
	"strconv"
)

var sexReg = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]*)</span></td>`)
var hightReg = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
var weightReg = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var carReg = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]*)</span></td>`)
var ageReg = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var salaryReg = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]*)</td>`)
var marryReg = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]*)</td>`)
var jiGuanReg = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]*)</td>`)
var xingZuoReg = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]*)</span></td>`)
var hotelReg = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]*)</span></td>`)
var occupationReg = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]*)</td>`)
var educationReg = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]*)</td>`)

func ParsePerson(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageReg))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, hightReg))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightReg))
	if err == nil {
		profile.Weight = weight
	}
	profile.Name = name
	profile.Marriage = string(extractString(contents, marryReg))
	profile.Gender = string(extractString(contents, sexReg))
	profile.Car = string(extractString(contents, carReg))
	profile.Hokou = string(extractString(contents, jiGuanReg))
	profile.House = string(extractString(contents, hotelReg))
	profile.Occupation = string(extractString(contents, occupationReg))
	profile.Education = string(extractString(contents, educationReg))
	profile.Xingzuo = string(extractString(contents, xingZuoReg))
	profile.Income = string(extractString(contents, salaryReg))

	result := engine.ParseResult{}
	result.Items = []interface{}{profile}

	return result
}
func extractString(contents []byte, reg *regexp.Regexp) string {
	match := reg.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
