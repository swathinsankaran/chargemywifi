package model

type Label string

const (
	LabelBatteryQuantity Label = "lDashBatteryQuantity"
	LabelChargeStatus    Label = "lDashChargeStatus"
)

const (
	JIOFIURL string = "http://jiofi.local.html/cgi-bin/en-jio/mStatus.html"
)