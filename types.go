package connectbox

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// GlobalSettings is a response format for getter.xml/fn=1 endpoint.
type GlobalSettings struct {
	AccessLevel               string `xml:"AccessLevel"`
	SwVersion                 string `xml:"SwVersion"`
	CmProvisionMode           string `xml:"CmProvisionMode"`
	DsLite                    string `xml:"DsLite"`
	GwProvisionMode           string `xml:"GwProvisionMode"`
	GwOperMode                string `xml:"GWOperMode"`
	ConfigVenderModel         string `xml:"ConfigVenderModel"`
	HideRemoteAccess          string `xml:"HideRemoteAccess"`
	HideModemMode             string `xml:"HideModemMode"`
	HideCustomerDHCPLANChange string `xml:"HideCustomerDhcpLanChange"`
	ShowDDNS                  string `xml:"ShowDDNS"`
	OperatorID                string `xml:"OperatorId"`
	AccessDenied              string `xml:"AccessDenied"`
	LockedOut                 string `xml:"LockedOut"`
	CountryID                 string `xml:"CountryID"`
	Title                     string `xml:"title"`
	Interface                 string `xml:"Interface"`
	OperStatus                string `xml:"operStatus"`
}

// CMSystemInfo is a response format for getter.xml/fn=2 endpoint.
type CMSystemInfo struct {
	DocsisMode      string `xml:"cm_docsis_mode"`
	HardwareVersion string `xml:"cm_hardware_version"`
	MacAddr         string `xml:"cm_mac_addr"`
	SerialNumber    string `xml:"cm_serial_number"`
	SystemUptime    int    `xml:"cm_system_uptime"`
	NetworkAccess   string `xml:"cm_network_access"`
}

// UnmarshalXML adds string to seconds conversion.
func (c *CMSystemInfo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias CMSystemInfo
	aux := &struct {
		*Alias
		SystemUptime string `xml:"cm_system_uptime"`
	}{
		Alias: (*Alias)(c),
	}

	if err := d.DecodeElement(&aux, &start); err != nil {
		return err //nolint:wrapcheck
	}

	dur, err := parseDuration(aux.SystemUptime)
	if err != nil {
		return err
	}
	c.SystemUptime = int(dur.Seconds())

	return nil
}

// Multilang is a response format for getter.xml/fn=3 endpoint.
type Multilang struct {
	WebCapPor string `xml:"WebCapPor"`
	Lang      string `xml:"Lang"`
}

// Status is a response format for getter.xml/fn=5 endpoint.
type Status struct {
	CMStatus             string `xml:"cm_status"`
	Bandmode             string `xml:"Bandmode"`
	BSSEnable2G          string `xml:"BssEnable2g"`
	SSID2G               string `xml:"SSID2G"`
	PreSharedKey2GLength string `xml:"PreSharedKey2gLength"`
	BssEnable5G          string `xml:"BssEnable5g"`
	SSID5G               string `xml:"SSID5G"`
	PreSharedKey5GLength string `xml:"PreSharedKey5gLength"`
	LANUserCount         string `xml:"LanUserCount"`
}

// Configuration is a response format for getter.xml/fn=6 endpoint.
type Configuration struct {
	FrequencyPlan string `xml:"FrequencyPlan"`
	Frequency     string `xml:"Frequency"`
}

// DownstreamTable is a response format for getter.xml/fn=10 endpoint.
type DownstreamTable struct {
	DsNum       string                      `xml:"ds_num"`
	Downstreams []DownstreamTableDownstream `xml:"downstream"`
}

// DownstreamTableDownstream is a part of DownstreamTable.
type DownstreamTableDownstream struct {
	Freq         string `xml:"freq"`
	Pow          string `xml:"pow"`
	Snr          string `xml:"snr"`
	Mod          string `xml:"mod"`
	Chid         string `xml:"chid"`
	RxMER        string `xml:"RxMER"`
	PreRs        string `xml:"PreRs"`
	PostRs       string `xml:"PostRs"`
	IsQamLocked  string `xml:"IsQamLocked"`
	IsFECLocked  string `xml:"IsFECLocked"`
	IsMpegLocked string `xml:"IsMpegLocked"`
}

// UpstreamTable is a response format for getter.xml/fn=11 endpoint.
type UpstreamTable struct {
	UsNum     string                  `xml:"us_num"`
	Upstreams []UpstreamTableUpstream `xml:"upstream"`
}

// UpstreamTableUpstream is a part of UpstreamTable.
type UpstreamTableUpstream struct {
	Usid        string `xml:"usid"`
	Freq        string `xml:"freq"`
	Power       string `xml:"power"`
	Srate       string `xml:"srate"`
	Mod         string `xml:"mod"`
	Ustype      string `xml:"ustype"`
	T1Timeouts  string `xml:"t1Timeouts"`
	T2Timeouts  string `xml:"t2Timeouts"`
	T3Timeouts  string `xml:"t3Timeouts"`
	T4Timeouts  string `xml:"t4Timeouts"`
	Channeltype string `xml:"channeltype"`
	MessageType string `xml:"messageType"`
}

// SignalTable is a response format for getter.xml/fn=12 endpoint.
type SignalTable struct {
	SigNum  string              `xml:"sig_num"`
	Signals []SignalTableSignal `xml:"signal"`
}

// SignalTableSignal is a part of SignalTable.
type SignalTableSignal struct {
	Dsid          string `xml:"dsid"`
	Unerrored     string `xml:"unerrored"`
	Correctable   string `xml:"correctable"`
	Uncorrectable string `xml:"uncorrectable"`
}

// EventLogTable is a response format for getter.xml/fn=13 endpoint.
type EventLogTable struct {
	EventLogs []EventLogTableEventLog `xml:"eventlog"`
}

// EventLogTableEventLog is a part of EventLogTable.
type EventLogTableEventLog struct {
	Prior string `xml:"prior"`
	Text  string `xml:"text"`
	Time  string `xml:"time"`
	T     string `xml:"t"`
}

// FirewallLogTable is a response format for getter.xml/fn=19 endpoint.
type FirewallLogTable struct {
	FirewallLogs []FirewallLogTableFirewallLog `xml:"firewalllog"`
}

// FirewallLogTableFirewallLog is a part of FirewallLogTable.
type FirewallLogTableFirewallLog struct {
	Prior string `xml:"prior"`
	Text  string `xml:"text"`
	Time  string `xml:"time"`
}

// Langsetlist is a response format for getter.xml/fn=21 endpoint.
type Langsetlist struct {
	LangSetSupport []string `xml:"langSet_support"`
}

// Fail is a response format for getter.xml/fn=22 endpoint.
type Fail struct {
	FailCount string `xml:"FailCount"`
}

// LoginTimer is a response format for getter.xml/fn=24 endpoint.
type LoginTimer struct {
	Flag        string `xml:"Flag"`
	AccessLevel string `xml:"AccessLevel"`
}

// LANSetting is a response format for getter.xml/fn=100 endpoint.
type LANSetting struct {
	UPnP             string `xml:"UPnP"`
	LANMAC           string `xml:"LanMAC"`
	LANIP            string `xml:"LanIP"`
	DMZAddr          string `xml:"DMZaddr"`
	DMZ              string `xml:"DMZ"`
	LanIPv6          string `xml:"LanIPv6"`
	LanIPv6Prefix    string `xml:"LanIPv6Prefix"`
	SubnetMask       string `xml:"subnetmask"`
	DHCPStartAddress string `xml:"DHCP_startaddress"`
	DHCPEndAddress   string `xml:"DHCP_endaddress"`
}

// DHCPv6Info is a response format for getter.xml/fn=103 endpoint.
type DHCPv6Info struct {
	AllowDHCPv6Setting          string `xml:"AllowDHCPv6Setting"`
	IPv6RAManagedFlag           string `xml:"ipv6RAManagedflag"`
	IPv6Saddr                   string `xml:"ipv6_saddr"`
	IPv6Prefix                  string `xml:"ipv6_prefix"`
	NumberOfAddr                string `xml:"NumberOfAddr"`
	IPv6PrefixPreferredLifeTime string `xml:"ipv6PrefixPreferredLifeTime"`
	IPv6PrefixValidLifeTime     string `xml:"ipv6PrefixValidLifeTime"`
	DHCPV6AddrLifeTime          string `xml:"dhcpV6AddrLifeTime"`
	IPv6RALifetime              string `xml:"ipv6RALifetime"`
	IPv6RAIntervaltime          string `xml:"ipv6RAIntervaltime"`
}

// BasicDHCP is a response format for getter.xml/fn=105 endpoint.
type BasicDHCP struct {
	EnableDHCPv4              string                    `xml:"enableDHCPv4"`
	AddrStart                 string                    `xml:"Addr_start"`
	NumberOfCpes              string                    `xml:"NumberOfCpes"`
	LeaseTime                 string                    `xml:"LeaseTime"`
	LanIP                     string                    `xml:"LanIP"`
	SubnetMask                string                    `xml:"subnetmask"`
	ReserveIPAddrs            []BasicDHCPReserveIPAddrs `xml:"ReserveIpadrr"`
	BlockSubnetIP             []string                  `xml:"BlockSubnetIP"`
	BlockSubnetMask           []string                  `xml:"BlockSubnetMask"`
	HideCustomerDHCPLANChange string                    `xml:"HideCustomerDhcpLanChange"`
}

// BasicDHCPReserveIPAddrs is a part of BasicDHCP.
type BasicDHCPReserveIPAddrs struct {
	MacAddress string `xml:"MacAddress"`
	LeasedIP   string `xml:"LeasedIP"`
}

// WANSetting is a response format for getter.xml/fn=107 endpoint.
type WANSetting struct {
	NAPTMode        string   `xml:"NAPT_mode"`
	WANMAC          string   `xml:"WanMAC"`
	WANIPv6Addrs    []string `xml:"wan_ipv6_addr>wan_ipv6_addr_entry"`
	WANDHCPv6Srv    string   `xml:"WanDhcpv6Srv"`
	IPv6LeaseTime   string   `xml:"ipv6_LeaseTime"`
	IPv6LeaseExpire string   `xml:"ipv6_LeaseExpire"`
	WANIPv6DNSAddr  []string `xml:"wan_ipv6_dnsaddr>wan_ipv6_dnsaddr_entry"`
	WANIP           string   `xml:"WanIP"`
	GatewayAddress  string   `xml:"gateway_address"`
	LeaseTime       string   `xml:"LeaseTime"`
	LeaseExpire     string   `xml:"LeaseExpire"`
	WANIPv4DNSAddr  []string `xml:"wan_ipv4_dnsaddr>wan_ipv4_dnsaddr_entry"`
	DsliteEnable    string   `xml:"dslite_enable"`
	DsliteFqdn      string   `xml:"dslite_fqdn"`
	DsliteAddr      string   `xml:"dslite_addr"`
}

// IPFiltering is a response format for getter.xml/fn=109 endpoint.
type IPFiltering struct {
	LanIP       string `xml:"LanIP"`
	SubnetMask  string `xml:"subnetmask"`
	TimeMode    string `xml:"time_mode"`
	GeneralTime string `xml:"GeneralTime"`
	DailyTime   string `xml:"DailyTime"`
}

// IPv6Filtering is a response format for getter.xml/fn=111 endpoint.
type IPv6Filtering struct {
	IPv6Prefix  string `xml:"ipv6_prefix"`
	Dir         string `xml:"dir"`
	TimeMode    string `xml:"time_mode"`
	GeneralTime string `xml:"GeneralTime"`
	DailyTime   string `xml:"DailyTime"`
}

// PortTrigger is a response format for getter.xml/fn=113 endpoint.
type PortTrigger struct{}

// WebFilter is a response format for getter.xml/fn=115 endpoint.
type WebFilter struct {
	FirewallProtection  string `xml:"firewallProtection"`
	BlockIPFragments    string `xml:"blockIpFragments"`
	PortScanDetection   string `xml:"portScanDetection"`
	SynFloodDetection   string `xml:"synFloodDetection"`
	ICMPFloodDetection  string `xml:"IcmpFloodDetection"`
	ICMPFloodDetectRate string `xml:"IcmpFloodDetectRate"`
}

// IPv6WebFilter is a response format for getter.xml/fn=117 endpoint.
type IPv6WebFilter struct {
	IPv6FirewallProtection  string `xml:"IPv6firewallProtection"`
	IPv6BlockIPFragments    string `xml:"IPv6blockIpFragments"`
	IPv6PortScanDetection   string `xml:"IPv6portScanDetection"`
	IPv6SynFloodDetection   string `xml:"IPv6synFloodDetection"`
	IPv6ICMPFloodDetection  string `xml:"IPv6IcmpFloodDetection"`
	IPv6ICMPFloodDetectRate string `xml:"IPv6IcmpFloodDetectRate"`
}

// MACFiltering is a response format for getter.xml/fn=119 endpoint.
type MACFiltering struct {
	MaxInstance string `xml:"maxInstance"`
	TimeMode    string `xml:"time_mode"`
	GeneralTime string `xml:"GeneralTime"`
	DailyTime   string `xml:"DailyTime"`
}

// Forwarding is a response format for getter.xml/fn=121 endpoint.
type Forwarding struct {
	LANIP      string           `xml:"LanIP"`
	SubnetMask string           `xml:"subnetmask"`
	UPnPs      []ForwardingUPnP `xml:"UPnP"`
}

// ForwardingUPnP is a part of Forwarding.
type ForwardingUPnP struct {
	LANIPAddr   string `xml:"LanIPAddr"`
	LANPort     string `xml:"LanPort"`
	WANPort     string `xml:"WanPort"`
	Protocol    string `xml:"Protocol"`
	Description string `xml:"Description"`
}

// LANUserTable is a response format for getter.xml/fn=123 endpoint.
type LANUserTable struct {
	Ethernet    []LANUserTableEthernet `xml:"Ethernet>clientinfo"`
	WIFI        []LANUserTableWIFI     `xml:"WIFI>clientinfo"`
	TotalClient string                 `xml:"totalClient"`
	Customer    string                 `xml:"Customer"`
}

// LANUserTableEthernet is a part of LANUserTable.
type LANUserTableEthernet struct {
	Interface   string `xml:"interface"`
	IPv4Addr    string `xml:"IPv4Addr"`
	XMLHostname string `xml:"xmlhostname"`
	XMLIcon     string `xml:"xmlicon"`
	Index       string `xml:"index"`
	InterfaceID string `xml:"interfaceid"`
	Hostname    string `xml:"hostname"`
	MACAddr     string `xml:"MACAddr"`
	Method      string `xml:"method"`
	LeaseTime   string `xml:"leaseTime"`
	Speed       string `xml:"speed"`
}

// LANUserTableWIFI is a part of LANUserTable.
type LANUserTableWIFI struct {
	Interface   string `xml:"interface"`
	IPv4Addr    string `xml:"IPv4Addr"`
	XMLHostname string `xml:"xmlhostname"`
	XMLIcon     string `xml:"xmlicon"`
	Index       string `xml:"index"`
	InterfaceID string `xml:"interfaceid"`
	Hostname    string `xml:"hostname"`
	MACAddr     string `xml:"MACAddr"`
	Method      string `xml:"method"`
	LeaseTime   string `xml:"leaseTime"`
	Speed       string `xml:"speed"`
}

// DDNS is a response format for getter.xml/fn=124 endpoint.
type DDNS struct {
	Enable       string `xml:"Enable"`
	DDNSProvider string `xml:"DDNSProvider"`
	Username     string `xml:"Username"`
	Password     string `xml:"Password"`
	Hostname     string `xml:"Hostname"`
	WanIP        string `xml:"WanIP"`
}

// RemoteAccess is a response format for getter.xml/fn=131 endpoint.
type RemoteAccess struct{}

// MTUSize is a response format for getter.xml/fn=134 endpoint.
type MTUSize struct {
	Size string `xml:"size"`
}

// CMState is a response format for getter.xml/fn=136 endpoint.
type CMState struct {
	TunnerTemperature int      `xml:"TunnerTemperature"`
	Temperature       int      `xml:"Temperature"`
	OperState         string   `xml:"OperState"`
	WANIPv4Addr       string   `xml:"wan_ipv4_addr"`
	WANIPv6Addrs      []string `xml:"wan_ipv6_addr>wan_ipv6_addr_entry"`
}

// UnmarshalXML adds fahrenheit to celsius conversion.
func (c *CMState) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias CMState
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err := d.DecodeElement(&aux, &start); err != nil {
		return err //nolint:wrapcheck
	}
	c.TunnerTemperature = fahrenheitToCelsius(c.TunnerTemperature)
	c.Temperature = fahrenheitToCelsius(c.Temperature)

	return nil
}

// WiredState1 is a response format for getter.xml/fn=137 endpoint.
type WiredState1 struct {
	Ports           []WiredState1Port `xml:"port"`
	Device          string            `xml:"Device"`
	EthFlaplistFile string            `xml:"ethflaplistFile"`
}

// WiredState1Port is a part of WiredState1.
type WiredState1Port struct {
	Eth   string `xml:"Eth"`
	Speed string `xml:"Speed"`
}

// WiredState2 is a response format for getter.xml/fn=143 endpoint.
type WiredState2 struct {
	Ports  []WiredState2Port `xml:"port"`
	Device string            `xml:"Device"`
}

// WiredState2Port is a part of WiredState2.
type WiredState2Port struct {
	Eth   string `xml:"Eth"`
	Speed string `xml:"Speed"`
}

// CMStatus is a response format for getter.xml/fn=144 endpoint.
type CMStatus struct {
	ProvisioningSt    string                `xml:"provisioning_st"`
	ProvisioningStNum string                `xml:"provisioning_st_num"`
	CMComment         string                `xml:"cm_comment"`
	DsNum             string                `xml:"ds_num"`
	Downstreams       []CMStatusDownstream  `xml:"downstream"`
	UsNum             string                `xml:"us_num"`
	Upstreams         []CMStatusUpstream    `xml:"upstream"`
	CMDocsisMode      string                `xml:"cm_docsis_mode"`
	CMNetworkAccess   string                `xml:"cm_network_access"`
	NumberOfCpes      string                `xml:"NumberOfCpes"`
	DMaxCpes          string                `xml:"dMaxCpes"`
	BpiEnable         string                `xml:"bpiEnable"`
	FileName          string                `xml:"FileName"`
	ServiceFlows      []CMStatusServiceFlow `xml:"serviceflow"`
}

// CMStatusDownstream is a part of CMStatus.
type CMStatusDownstream struct {
	Freq            string `xml:"freq"`
	Mod             string `xml:"mod"`
	Chid            string `xml:"chid"`
	State           string `xml:"state"`
	Status          string `xml:"status"`
	PrimarySettings string `xml:"primarySettings"`
}

// CMStatusUpstream is a part of CMStatus.
type CMStatusUpstream struct {
	Usid  string `xml:"usid"`
	Freq  string `xml:"freq"`
	Power string `xml:"power"`
	Srate string `xml:"srate"`
	State string `xml:"state"`
}

// CMStatusServiceFlow is a part of CMStatus.
type CMStatusServiceFlow struct {
	Sfid             string `xml:"Sfid"`
	Direction        string `xml:"direction"`
	PMaxTrafficRate  string `xml:"pMaxTrafficRate"`
	PMaxTrafficBurst string `xml:"pMaxTrafficBurst"`
	PMinReservedRate string `xml:"pMinReservedRate"`
	PMaxConcatBurst  string `xml:"pMaxConcatBurst"`
	PSchedulingType  string `xml:"pSchedulingType"`
}

// EthFlaplist is a response format for getter.xml/fn=147 endpoint.
type EthFlaplist struct {
	EthFlaplistFile string `xml:"ethflaplistFile"`
}

// WirelessBasic1 is a response format for getter.xml/fn=300 endpoint.
type WirelessBasic1 struct {
	NvCountry            string `xml:"NvCountry"`
	Bandmode             string `xml:"Bandmode"`
	ChannelRange         string `xml:"ChannelRange"`
	BSSEnable2G          string `xml:"BssEnable2g"`
	SSID2G               string `xml:"SSID2G"`
	HideNetwork2G        string `xml:"HideNetwork2G"`
	BandWidth2G          string `xml:"BandWidth2G"`
	BSSCoexistence       string `xml:"BssCoexistence"`
	TransmissionRate2G   string `xml:"TransmissionRate2g"`
	TransmissionMode2G   string `xml:"TransmissionMode2g"`
	SecurityMode2G       string `xml:"SecurityMode2g"`
	MulticastRate2G      string `xml:"MulticastRate2G"`
	ChannelSetting2G     string `xml:"ChannelSetting2G"`
	CurrentChannel2G     string `xml:"CurrentChannel2G"`
	PreSharedKey2G       string `xml:"PreSharedKey2g"`
	GroupRekeyInterval2G string `xml:"GroupRekeyInterval2g"`
	WpaAlgorithm2G       string `xml:"WpaAlgorithm2G"`
	SONAdminStatus       string `xml:"SONAdminStatus"`
	SONOperationalStatus string `xml:"SONOperationalStatus"`
	BssEnable5G          string `xml:"BssEnable5g"`
	SSID5G               string `xml:"SSID5G"`
	HideNetwork5G        string `xml:"HideNetwork5G"`
	BandWidth5G          string `xml:"BandWidth5G"`
	TransmissionRate5G   string `xml:"TransmissionRate5g"`
	TransmissionMode5G   string `xml:"TransmissionMode5g"`
	SecurityMode5G       string `xml:"SecurityMode5g"`
	MulticastRate5G      string `xml:"MulticastRate5G"`
	ChannelSetting5G     string `xml:"ChannelSetting5G"`
	CurrentChannel5G     string `xml:"CurrentChannel5G"`
	PreSharedKey5G       string `xml:"PreSharedKey5g"`
	GroupRekeyInterval5G string `xml:"GroupRekeyInterval5g"`
	WpaAlgorithm5G       string `xml:"WpaAlgorithm5G"`
}

// WirelessWmm is a response format for getter.xml/fn=302 endpoint.
type WirelessWmm struct {
	WMM2G              string `xml:"WMM2G"`
	Apsd2G             string `xml:"Apsd2G"`
	TransmissionMode2G string `xml:"TransmissionMode2g"`
	WMM5G              string `xml:"WMM5G"`
	Apsd5G             string `xml:"Apsd5G"`
	TransmissionMode5G string `xml:"TransmissionMode5g"`
}

// WirelessSiteSurvey is a response format for getter.xml/fn=305 endpoint.
type WirelessSiteSurvey struct {
	Count2G     string `xml:"count2G"`
	Count5G     string `xml:"count5G"`
	BandMode24G string `xml:"BandMode_2_4G"`
	BandMode5G  string `xml:"BandMode_5G"`
}

// WirelessGuestNetwork1 is a response format for getter.xml/fn=307 endpoint.
type WirelessGuestNetwork1 struct {
	MainEnable2G string                             `xml:"MainEnable2G"`
	MainEnable5G string                             `xml:"MainEnable5G"`
	Interfaces   []WirelessGuestNetwork1Interface   `xml:"Interface"`
	Interfaces5G []WirelessGuestNetwork1Interface5G `xml:"Interface5G"`
}

// WirelessGuestNetwork1Interface is a part of WirelessGuestNetwork1.
type WirelessGuestNetwork1Interface struct {
	Enable2G             string `xml:"Enable2G"`
	BSSID2G              string `xml:"BSSID2G"`
	GuestMac2G           string `xml:"GuestMac2G"`
	HideNetwork2G        string `xml:"HideNetwork2G"`
	SecurityMode2G       string `xml:"SecurityMode2g"`
	PreSharedKey2G       string `xml:"PreSharedKey2g"`
	GroupRekeyInterval2G string `xml:"GroupRekeyInterval2g"`
	WPAAlgorithm2G       string `xml:"WpaAlgorithm2G"`
}

// WirelessGuestNetwork1Interface5G is a part of WirelessGuestNetwork1.
type WirelessGuestNetwork1Interface5G struct {
	Enable5G             string `xml:"Enable5G"`
	BSSID5G              string `xml:"BSSID5G"`
	GuestMac5G           string `xml:"GuestMac5G"`
	HideNetwork5G        string `xml:"HideNetwork5G"`
	SecurityMode5G       string `xml:"SecurityMode5g"`
	PreSharedKey5G       string `xml:"PreSharedKey5g"`
	GroupRekeyInterval5G string `xml:"GroupRekeyInterval5g"`
	WPAAlgorithm5G       string `xml:"WpaAlgorithm5G"`
}

// CMWirelessWPS1 is a response format for getter.xml/fn=309 endpoint.
type CMWirelessWPS1 struct {
	MainEnable2G   string `xml:"MainEnable2g"`
	MainEnable5G   string `xml:"MainEnable5g"`
	WPSEnable24G   string `xml:"WpsEnable24G"`
	WPSEnable5G    string `xml:"WpsEnable5G"`
	WPSMethod24G   string `xml:"WpsMethod24G"`
	WPSMethod5G    string `xml:"WpsMethod5G"`
	WPSAPPin24G    string `xml:"WpsAPPIN24G"`
	WPSAPPin5G     string `xml:"WpsAPPIN5G"`
	WPSPinNum24G   string `xml:"WpsPINNUM24G"`
	WPSPinNum5G    string `xml:"WpsPINNUM5G"`
	WPSEnablePBC   string `xml:"WpsEnablePBC"`
	WPSEnablePIN   string `xml:"WpsEnablePIN"`
	WPSEnablePBC5G string `xml:"WpsEnablePBC5G"`
	WPSEnablePIN5G string `xml:"WpsEnablePIN5G"`
}

// CMWirelessAccessControl is a response format for getter.xml/fn=311 endpoint.
type CMWirelessAccessControl struct {
	BandMode           string                                    `xml:"BandMode"`
	BSSEnable2G        string                                    `xml:"BssEnable2g"`
	BSSEnable5G        string                                    `xml:"BssEnable5g"`
	SSID2G             string                                    `xml:"SSID2G"`
	SSID5G             string                                    `xml:"SSID5G"`
	HideNetwork2G      string                                    `xml:"HideNetwork2G"`
	HideNetwork5G      string                                    `xml:"HideNetwork5G"`
	SecurityMode2G     string                                    `xml:"SecurityMode2g"`
	SecurityMode5G     string                                    `xml:"SecurityMode5g"`
	PreSharedKey2G     string                                    `xml:"PreSharedKey2g"`
	PreSharedKey5G     string                                    `xml:"PreSharedKey5g"`
	WpaAlgorithm2G     string                                    `xml:"WpaAlgorithm2G"`
	WpaAlgorithm5G     string                                    `xml:"WpaAlgorithm5G"`
	AccessMode24G      string                                    `xml:"AccessMode24G"`
	AccessMode5G       string                                    `xml:"AccessMode5G"`
	BSSAccessEntries   []CMWirelessAccessControlBSSAccessEntry   `xml:"BssAccessEntry"`
	BSSAccessEntries5G []CMWirelessAccessControlBSSAccessEntry5G `xml:"BssAccessEntry5G"`
}

// CMWirelessAccessControlBSSAccessEntry is a part of CMWirelessAccessControl.
type CMWirelessAccessControlBSSAccessEntry struct {
	AccessStation    string `xml:"AccessStation"`
	AccessDeviceName string `xml:"AccessDeviceName"`
}

// CMWirelessAccessControlBSSAccessEntry5G is a part of CMWirelessAccessControl.
type CMWirelessAccessControlBSSAccessEntry5G struct {
	AccessStation5G    string `xml:"AccessStation5G"`
	AccessDeviceName5G string `xml:"AccessDeviceName5G"`
}

// ChannelMap is a response format for getter.xml/fn=313 endpoint.
type ChannelMap struct {
	Count2G            string                `xml:"count2G"`
	MyCurrentChannel2G string                `xml:"MyCurrentChannel2G"`
	Count5G            string                `xml:"count5G"`
	MyCurrentChannel5G string                `xml:"MyCurrentChannel5G"`
	BandMode24G        ChannelMapBandMode24G `xml:"BandMode_2_4G"`
	BandMode5G         ChannelMapBandMode5G  `xml:"BandMode_5G"`
}

// ChannelMapBandMode24G is a part of ChannelMap.
type ChannelMapBandMode24G struct {
	W2GCH1    string `xml:"W2GCH1"`
	W2GCH2    string `xml:"W2GCH2"`
	W2GCH3    string `xml:"W2GCH3"`
	W2GCH4    string `xml:"W2GCH4"`
	W2GCH5    string `xml:"W2GCH5"`
	W2GCH6    string `xml:"W2GCH6"`
	W2GCH7    string `xml:"W2GCH7"`
	W2GCH8    string `xml:"W2GCH8"`
	W2GCH9    string `xml:"W2GCH9"`
	W2GCH10   string `xml:"W2GCH10"`
	W2GCH11   string `xml:"W2GCH11"`
	W2GCH12   string `xml:"W2GCH12"`
	W2GCH13   string `xml:"W2GCH13"`
	Maxaxis2G string `xml:"maxaxis2G"`
	Total2G   string `xml:"total2g"`
}

// ChannelMapBandMode5G is a part of ChannelMap.
type ChannelMapBandMode5G struct {
	W5GCH1    string `xml:"W5GCH1"`
	W5GCH2    string `xml:"W5GCH2"`
	W5GCH3    string `xml:"W5GCH3"`
	W5GCH4    string `xml:"W5GCH4"`
	W5GCH5    string `xml:"W5GCH5"`
	W5GCH6    string `xml:"W5GCH6"`
	W5GCH7    string `xml:"W5GCH7"`
	W5GCH8    string `xml:"W5GCH8"`
	W5GCH9    string `xml:"W5GCH9"`
	W5GCH10   string `xml:"W5GCH10"`
	W5GCH11   string `xml:"W5GCH11"`
	W5GCH12   string `xml:"W5GCH12"`
	W5GCH13   string `xml:"W5GCH13"`
	W5GCH14   string `xml:"W5GCH14"`
	W5GCH15   string `xml:"W5GCH15"`
	W5GCH16   string `xml:"W5GCH16"`
	W5GCH17   string `xml:"W5GCH17"`
	W5GCH18   string `xml:"W5GCH18"`
	W5GCH19   string `xml:"W5GCH19"`
	Maxaxis5G string `xml:"maxaxis5G"`
	Total5G   string `xml:"total5g"`
}

// WirelessBasic2 is a response format for getter.xml/fn=315 endpoint.
type WirelessBasic2 struct {
	Bandmode       string `xml:"Bandmode"`
	BSSEnable2G    string `xml:"BssEnable2g"`
	BSSEnable5G    string `xml:"BssEnable5g"`
	WiFiChipStatus string `xml:"WiFi_chip_status"`
	CMStatus       string `xml:"cm_status"`
}

// WirelessGuestNetwork2 is a response format for getter.xml/fn=317 endpoint.
type WirelessGuestNetwork2 struct {
	Year        string                           `xml:"year"`
	Mouth       string                           `xml:"mouth"`
	Day         string                           `xml:"day"`
	Hour        string                           `xml:"hour"`
	Minute      string                           `xml:"minute"`
	Interface   WirelessGuestNetwork2Interface   `xml:"Interface"`
	Interface5G WirelessGuestNetwork2Interface5G `xml:"Interface5G"`
}

// WirelessGuestNetwork2Interface is a part of WirelessGuestNetwork2.
type WirelessGuestNetwork2Interface struct {
	MainEnable2G         string `xml:"MainEnable2G"`
	Enable2G             string `xml:"Enable2G"`
	BSSID2G              string `xml:"BSSID2G"`
	GuestMac2G           string `xml:"GuestMac2G"`
	HideNetwork2G        string `xml:"HideNetwork2G"`
	SecurityMode2G       string `xml:"SecurityMode2g"`
	PreSharedKey2G       string `xml:"PreSharedKey2g"`
	GroupRekeyInterval2G string `xml:"GroupRekeyInterval2g"`
	WPAAlgorithm2G       string `xml:"WpaAlgorithm2G"`
}

// WirelessGuestNetwork2Interface5G is a part of WirelessGuestNetwork2.
type WirelessGuestNetwork2Interface5G struct {
	MainEnable5G         string `xml:"MainEnable5G"`
	Enable5G             string `xml:"Enable5G"`
	BSSID5G              string `xml:"BSSID5G"`
	GuestMac5G           string `xml:"GuestMac5G"`
	HideNetwork5G        string `xml:"HideNetwork5G"`
	SecurityMode5G       string `xml:"SecurityMode5g"`
	PreSharedKey5G       string `xml:"PreSharedKey5g"`
	GroupRekeyInterval5G string `xml:"GroupRekeyInterval5g"`
	WPAAlgorithm5G       string `xml:"WpaAlgorithm5G"`
}

// WirelessClient is a response format for getter.xml/fn=322 endpoint.
type WirelessClient struct {
	Client2G []WirelessClientClient2G `xml:"Client2G"`
	Client5G []WirelessClientClient5G `xml:"Client5G"`
}

// WirelessClientClient2G is a part of WirelessClient.
type WirelessClientClient2G struct {
	ClientInfo []WirelessClientClient2GClientInfo `xml:"clientinfo"`
}

// WirelessClientClient5G is a part of WirelessClient.
type WirelessClientClient5G struct {
	ClientInfo []WirelessClientClient5GClientInfo `xml:"clientinfo"`
}

// WirelessClientClient2GClientInfo is a part of WirelessClientClient2G.
type WirelessClientClient2GClientInfo struct {
	SSID          string `xml:"SSID"`
	MAC           string `xml:"MAC"`
	PhyRateTx     string `xml:"phy_rate_tx"`
	PhyRateRx     string `xml:"phy_rate_rx"`
	PhyMode       string `xml:"phy_mode"`
	AuthMode      string `xml:"Auth_mode"`
	RSSI          string `xml:"RSSI"`
	EncryptMethod string `xml:"EncryptMethod"`
}

// WirelessClientClient5GClientInfo is a part of WirelessClientClient5G.
type WirelessClientClient5GClientInfo struct {
	SSID          string `xml:"SSID"`
	MAC           string `xml:"MAC"`
	PhyRateTx     string `xml:"phy_rate_tx"`
	PhyRateRx     string `xml:"phy_rate_rx"`
	PhyMode       string `xml:"phy_mode"`
	AuthMode      string `xml:"Auth_mode"`
	RSSI          string `xml:"RSSI"`
	EncryptMethod string `xml:"EncryptMethod"`
}

// CMWirelessWPS2 is a response format for getter.xml/fn=323 endpoint.
type CMWirelessWPS2 struct {
	WPSStat   string `xml:"WPS_stat"`
	WPSResult string `xml:"WPS_result"`
}

// DefaultValue is a response format for getter.xml/fn=324 endpoint.
type DefaultValue struct {
	LoginPwd string `xml:"loginPwd"`
	WIFISSID string `xml:"WiFiSSID"`
	WIFIkey  string `xml:"WiFikey"`
}

// GstRandomPassword is a response format for getter.xml/fn=325 endpoint.
type GstRandomPassword struct {
	PreSharedKey string `xml:"PreSharedKey"`
}

// WIFIState is a response format for getter.xml/fn=326 endpoint.
type WIFIState struct {
	Primary24G string `xml:"primary24g"`
	Primary5G  string `xml:"primary5g"`
}

// WirelessResetting is a response format for getter.xml/fn=328 endpoint.
type WirelessResetting struct {
	IsWirelessResetting string `xml:"isWirelessResetting"`
}

var durationRegexp = regexp.MustCompile(`(?:(\d+)day\(s\))?(\d+)h:(\d+)m:(\d+)s`)

// Input format: "1day(s)2h:34m:56s".
func parseDuration(s string) (time.Duration, error) {
	matches := durationRegexp.FindStringSubmatch(s)
	if len(matches) != 5 {
		return 0, fmt.Errorf("invalid duration string")
	}

	days, _ := strconv.Atoi(matches[1])
	hours, _ := strconv.Atoi(matches[2])
	minutes, _ := strconv.Atoi(matches[3])
	seconds, _ := strconv.Atoi(matches[4])

	dur := time.Duration(days)*24*time.Hour +
		time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second

	return dur, nil
}

func fahrenheitToCelsius(f int) int {
	return (f - 32) * 5.0 / 9
}
