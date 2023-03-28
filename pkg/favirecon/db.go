/*
favirecon - Use favicon.ico to improve your target recon phase. Quickly detect technologies, WAF, exposed panels, known services.

This repository is under MIT License https://github.com/edoardottt/favirecon/blob/main/LICENSE
*/

package favirecon

import (
	"errors"
	"fmt"

	"github.com/projectdiscovery/goflags"
)

// nolint: gochecknoglobals
var (
	db = map[string]string{
		"-1000719429": "SpamSniper",
		"-1001050714": "GROWI - Team collaboration software",
		"-1003107038": "WordPress",
		"1004209915":  "Contaware",
		"-1004700569": "TheHost",
		"-1010568750": "phpMyAdmin",
		"-1012744956": "Kaspersky Endpoint Security Cloud",
		"1012948051":  "Opengear Management Console",
		"1015545776":  "pfSense",
		"-1015932800": "Ghost (CMS)",
		"1016158463":  "Netonix",
		"-101718582":  "GIAE Online",
		"-1017346469": "WeBWork",
		"1020814938":  "Ubiquiti - AirOS",
		"1021528395":  "Phicomm",
		"-1022206565": "CrushFTP",
		"-1022319330": "Belkin router",
		"1023869165":  "Flood Web UI",
		"1027252292":  "AnyDesk",
		"-1028703177": "TP Link",
		"1029363268":  "QSolve Management System",
		"-103091179":  "Halon Security",
		"1032553729":  "Elastix",
		"-1036244016": "XWeb",
		"1037387972":  "Dlink Router",
		"-103848262":  "Makhost",
		"1038500535":  "D-Link (router/network)",
		"-1038557304": "Webmin",
		"1039044672":  "Apache",
		"-1041180225": "QNAP NAS Virtualization Station",
		"104189364":   "Vigor Router",
		"-104250814":  "idfuse",
		"-1043355641": "eBuilder",
		"1045696447":  "Sophos User Portal/VPN Portal",
		"1046701084":  "iCanteen",
		"-10469226":   "3CX Web",
		"1047213685":  "Netgear (Network)",
		"-1049382096": "BrbOS",
		"-1050377603": "ThingsPro Gateway",
		"-1050786453": "Plesk",
		"105083909":   "FireEye",
		"-1051063062": "Apache Archiva",
		"-1051205943": "Siemens Simatic",
		"1051648103":  "Securepoint",
		"1053515370":  "OfficeSpace",
		"-1054477011": "Leica Geosystems",
		"1056758035":  "Webmin",
		"-1057103077": "CPT",
		"-1057698644": "Issabel",
		"1059329877":  "Tecvoz",
		"1059461494":  "Ecessa WAN Optimizer",
		"-1059717365": "Ebay",
		"-1060318941": "WayOS",
		"-1062191521": "VOS",
		"1064742722":  "RabbitMQ",
		"-106646451":  "WISPR (Airlan)",
		"-1067420240": "GraphQL",
		"1068026086":  "ComfortClick",
		"-1068237086": "Etherpad",
		"106844876":   "Revive Adserver",
		"107061220":   "T-Mobile",
		"-1071169404": "Deluge Web UI",
		"1073075415":  "Jaspersoft",
		"1075305934":  "IBM Lotus Domino Web Access",
		"1076320928":  "XOffice",
		"1076833462":  "Mailu",
		"1081719753":  "D-Link (Network)",
		"-1082845244": "RabbitMQ Management",
		"108362015":   "IMuse",
		"108411803":   "Deluge Web UI",
		"-108457676":  "QMatic",
		"-1085284672": "Wordpress",
		"1088281959":  "RF-301K",
		"-1088664572": "Microbit",
		"1090061843":  "Webtitan Cloud",
		"-1093172228": "truVision (NVR)",
		"-1093314638": "IBM iNotes",
		"1095915848":  "Airwatch",
		"-1096293925": "ECO World UI",
		"-10974981":   "Shinobi (CCTV)",
		"1098596251":  "Fedora Server Edition",
		"1101614703":  "UPS",
		"1102047973":  "KPMG",
		"1103599349":  "Untangle",
		"-1104274084": "eRad",
		"1104879639":  "MandeUmZap",
		"110768013":   "Sapido router",
		"-1108647419": "SmartFoxServer",
		"-1109801483": "Seafile",
		"-1111841537": "Platform For Science",
		"-111232642":  "starbeacon CMS",
		"-1115903764": "Pydio",
		"1117165781":  "SimpleHelp (Remote Support)",
		"-1118274094": "Communigate",
		"1118684072":  "Baidu",
		"-1119613926": "Bluehost",
		"-1124868062": "Netport Software (DSL)",
		"1122883478":  "Codiad",
		"1124727888":  "Powered by Leapfunder",
		"112509862":   "i.LON",
		"1125166869":  "KEMP Login Screen",
		"-1125259127": "Avigilon Web Access",
		"-1126048407": "ParentSchool SIS",
		"1126652438":  "Cypress Configuration and Management",
		"1126835021":  "InstaCart",
		"-112697680":  "PostGraphile",
		"-1128128488": "VMWare Access Gateway",
		"-1130575070": "WISE TIGER router",
		"-1131046522": "Dolibarr",
		"1135165421":  "Ricoh",
		"-1131689409": "yTM",
		"-1132465110": "NiFi server",
		"1135750934":  "Infinite Campus",
		"1136074368":  "SchoolNet",
		"-1138074337": "LUM",
		"-1139576724": "WorkForce Telestaff Login",
		"1139788073":  "Metasploit",
		"1142227528":  "Aruba (Virtual Controller)",
		"-1142667323": "Tenderland",
		"1143877728":  "Dremio",
		"1144925962":  "Dlink Webcam",
		"1147858285":  "Dell",
		"-1148190371": "OPNsense",
		"-1148627124": "Nagios XI",
		"1149842582":  "Improbable.io",
		"-1151246177": "Servatica",
		"-1151675028": "ISP Manager (Web Hosting Panel)",
		"-1153873472": "Airwatch",
		"-1153950306": "Dell",
		"-1155883524": "Homebridge",
		"1156815778":  "JBoss EAP 6",
		"1157789622":  "Ubiquiti UNMS",
		"-1160966609": "Gitea",
		"1162358134":  "ALLNET",
		"-1162630024": "Jumpserver",
		"-1162730477": "Vanderbilt SPC",
		"116323821":   "Spring boot",
		"1163764264":  "CentreStack",
		"1165640109":  "Mizuho Global e-Banking",
		"1165838194":  "ThinkPHP",
		"1166865883":  "Haivision Makito",
		"1167359424":  "Samsung",
		"1169183049":  "BoaServer",
		"-1169314298": "INSTAR IP Cameras",
		"1170430551":  "React App",
		"-1170501093": "phpLDAPadmin",
		"-1171707951": "Accellion File Transfer",
		"-1172436532": "SecureLink Enterprise Access",
		"1174526129":  "OLT Web management interface",
		"1174841451":  "Drupal",
		"1175383902":  "Real 2D/3D",
		"1176554158":  "Sauter moduWeb",
		"-1176564322": "HWg-STE",
		"1176619560":  "Intelbras",
		"1177920756":  "AVS electronics",
		"1178471380":  "Foldr",
		"1179099333":  "Youtube",
		"1180765435":  "OrientDB",
		"1182206475":  "VMWare Horizon",
		"-1182250791": "Kodi Web Interface",
		"-1183519486": "EMail Marketing Service",
		"-1184167764": "DEVA Broadcast",
		"118652228":   "ZF2 App",
		"1188424430":  "KT&C",
		"1188645141":  "Huweishen",
		"-1188918488": "Fujitsu NAS",
		"-1189561879": "Linear eMerge",
		"1192261743":  "E2open",
		"-1193783166": "HOT Box",
		"1194900831":  "boost ai",
		"119532476":   "Mahara ePortfolio system",
		"-1196651128": "NetGear Login",
		"1197632109":  "Video Conferencing server for Enterprise",
		"119741608":   "Teltonika",
		"-1197966750": "Gemalto",
		"12003995":    "Walmart",
		"-1200737715": "Kibana",
		"-1202103423": "TeamSystem",
		"-1203021870": "Kubeflow",
		"-1205024243": "lwIP (A Lightweight TCP/IP stack)",
		"1205288142":  "OpenVidu",
		"-120854511":  "CUPS",
		"1211608009":  "Openfire Admin Console",
		"-1212471026": "PBX in a flash",
		"-1214233731": "Wightman",
		"1215779410":  "ControlID",
		"-1217002480": "AppDynamics",
		"1221759509":  "Dlink Webcam",
		"-121875991":  "Xen Orchestra",
		"1224108256":  "CrownPeak",
		"1224439683":  "Barracuda Networks",
		"1224994964":  "Vodafone",
		"-1225484776": "Endian Firewall",
		"-1226343463": "SnapGear Management Console",
		"1227004061":  "Assurant",
		"1227052603":  "Alibaba Cloud (Block Page)",
		"1227152965":  "Killing Floor Web Admin",
		"1228349756":  "h5ai",
		"-1231308448": "Atlassian Crowd",
		"-1231681737": "Ghost (CMS)",
		"1232159009":  "Apple",
		"1232596212":  "OpenStack",
		"1232683746":  "Citigold",
		"1233473326":  "AmicaWEB - Combivox",
		"1234311970":  "Onera",
		"1235070469":  "Synology VPN Plus",
		"-1235192469": "Metasploit",
		"-1237480336": "Vue",
		"123821839":   "Sangfor",
		"-1238307632": "Devline",
		"-1238837624": "Liquid Pixels",
		"1239721344":  "Micro Focus Vibe",
		"-1240222446": "Zhejiang Uniview Technologies Co.",
		"-1240387895": "Unqork",
		"1240573703":  "Roche Global Corporate Website",
		"1240618871":  "Oracle Hospitality - Reporting and Analytics",
		"1241049726":  "iomega NAS",
		"1241305867":  "Enlighted Manage",
		"-124234705":  "VoIPmonitor",
		"-1243674871": "AppNext",
		"1244636413":  "cPanel Login",
		"-1244757483": "Kaspersky",
		"1245897666":  "xWiki",
		"1248917303":  "JupyterHub",
		"1249285083":  "Ubiquiti Aircube",
		"-1249852061": "Microsoft Outlook",
		"-1250210868": "bticino",
		"-1250474341": "VMware Workspace",
		"-1250820636": "CP Plus",
		"1251810433":  "Cafe24 (Korea)",
		"-1252041730": "Vue.js",
		"1253536875":  "ZNC Web Frontend",
		"-1253943670": "HUMAX",
		"-1255454840": "Mailwatch",
		"-1255347784": "AngularJS",
		"-1255992602": "VMware Horizon",
		"-1258058404": "TileServer GL",
		"-1259700605": "Forcepoint Access forbidden",
		"1262005940":  "Jamf Pro Login",
		"-1262113920": "Git",
		"1263391110":  "Web Directory",
		"1266034062":  "Mailing lists service",
		"-1267819858": "KeyHelp (Keyweb AG)",
		"-1268095485": "VZPP Plesk",
		"-12700016":   "Seafile",
		"-1272756243": "Citrix Login",
		"-1274798165": "Sophos Mobile",
		"1273982002":  "Mautic (Open Source Marketing Automation)",
		"1274078387":  "TP-LINK (Network Device)",
		"-1275148624": "Accrisoft",
		"-1275226814": "XAMPP",
		"-1277814690": "LaCie",
		"1278323681":  "Gitlab",
		"-127886975":  "Metasploit",
		"1280124172":  "VK Services",
		"1280461262":  "Braze",
		"1280907310":  "Webmin",
		"1281253102":  "Dahua Storm (DVR)",
		"1288385475":  "Livebox",
		"1292731542":  "IP-Symcon WebFront",
		"129457226":   "Liferay Portal",
		"-1298108480": "Yii PHP Framework",
		"1298613326":  "NEXIOS",
		"1302486561":  "NetData",
		"1303364595":  "ErpNext",
		"1307188176":  "UniFi NVR: Software Portal",
		"1307375944":  "Octoprint (3D printer)",
		"-1308101551": "Pegasus",
		"130960039":   "EnergyTeam",
		"1312262473":  "MicroEngine",
		"-1312806261": "Confluence",
		"-1313140993": "vBulletin",
		"-131381065":  "eve - Emulated Virtual Environment",
		"-1314102215": "Verdaccio",
		"-1314405728": "Alexandria",
		"-1314943136": "Apache HTTP Server",
		"-1316071138": "TalkingData",
		"-1316507119": "Tieline",
		"1318124267":  "Avigilon",
		"-1319025408": "Netgear",
		"-1319399032": "ispbox",
		"1319699698":  "Form.io",
		"1319780356":  "Biotime",
		"1323738809":  "Nexus Repository Manager",
		"1325605650":  "Overleaf",
		"133039858":   "Libraesva ESG",
		"1332183546":  "Sun Java System Application Server",
		"-1332828203": "Network Storage Server",
		"1333537166":  "Alfresco",
		"-1334408578": "Solar-Log",
		"1334507537":  "BroadWave",
		"-1334617766": "Jitsi Meet",
		"-1335251146": "Totolink",
		"1335824931":  "AVer",
		"1335852856":  "RedCap",
		"-1336042616": "S2 Netbox Login",
		"-1338133217": "Riverbed",
		"-1343070146": "Intelbras SA",
		"-1343761772": "SAVPizzaDoor",
		"-134375033":  "Plesk",
		"-134458680":  "LifeSize UVC Video Center",
		"-1344736688": "Phicomm",
		"1346408936":  "PowerController",
		"-1346447358": "TilginAB (HomeGateway)",
		"1347937389":  "SAP Conversational AI",
		"-1348984223": "Jaeger UI",
		"1349199107":  "BlueBean",
		"-1350437236": "WiFi Login",
		"-1350587650": "nxFilter",
		"1351214488":  "Dead End",
		"-1351901211": "Luma Surveillance",
		"-1352019987": "Viprinet",
		"-1353325588": "Remote UI Portal",
		"-1354933624": "Dlink Webcam",
		"-1355675073": "RadiusManager",
		"1356662359":  "Outlook Web Application",
		"1356900853":  "Winmail Server",
		"135741046":   "FG Forrest",
		"1358855492":  "PrivateBin",
		"-1360098475": "WorkWave",
		"-1364880648": "InnoMedia",
		"1366863347":  "Greenshift",
		"1367373012":  "DolphinPHP",
		"1367958180":  "litemall",
		"1370528867":  "Yahoo",
		"1370990243":  "Trend Micro Deep Security Manager",
		"-137295400":  "NETGEAR ReadyNAS",
		"-1373888521": "ErpNext",
		"-1374555452": "Baidu",
		"1375401192":  "VICIdial",
		"1375487984":  "Web based Configurator",
		"-1375671634": "The Lounge",
		"-1376750468": "Greatek",
		"1376888327":  "Ecole",
		"-1378182799": "Archivematica",
		"-1379982221": "Atlassian - Bamboo",
		"1380259394":  "BBC",
		"1381079566":  "EVSE Web Interface",
		"-1381126564": "Global Management System",
		"1382324298":  "Apple",
		"-1385513933": "WangShen",
		"-1387041300": "PINC Solutions",
		"-1387795012": "PayKeeper",
		"1390393078":  "Optergy",
		"-1393435578": "Mediatrix",
		"1395038954":  "Phenix",
		"-1395400951": "Huawei - ADSL/Router",
		"1398098445":  "ownCloud",
		"-1399433489": "Prometheus Server",
		"1403071546":  "Pearson",
		"1405460984":  "pfSense",
		"1407809463":  "Dolibarr",
		"1410610129":  "Supermicro Intelligent Management (IPMI)",
		"-1413670889": "Ansible AWX",
		"-1414475558": "Microsoft IIS",
		"-1416401235": "Gigapod",
		"1417992821":  "Lynx System",
		"-1418620977": "Wekan",
		"142044466":   "The Internet Archive",
		"1422770566":  "Endian Firewall",
		"142313513":   "Facebook",
		"1423291318":  "Smart PABX",
		"1423930703":  "lighttpd",
		"-1424036600": "Portainer",
		"-1424697070": "SmartAdServer",
		"1424295654":  "Icecast Streaming Media Server",
		"1427774978":  "Home motion by Somfy",
		"1427976651":  "ZTE (Network)",
		"-1432009195": "Elastix",
		"1433417005":  "Salesforce",
		"1435180442":  "HOTSPLOTS Router",
		"-1435467015": "CoreLogic",
		"1436966696":  "Barracuda",
		"-1437701105": "XAMPP",
		"-1439222863": "Ivanti",
		"-1441715561": "Vodafone",
		"-1441956789": "Tableau",
		"1442699674":  "NET",
		"-1442789563": "Nuxt JS",
		"1443361300":  "Poweradmin",
		"-1445519482": "Kaseya",
		"1446490141":  "MetInfo",
		"1446754233":  "Apache",
		"-1446794564": "Ubiquiti Login Portals",
		"-1450898105": "TELE2",
		"-1451027808": "Intelbras",
		"1451166122":  "ExtremeXOS ScreenPlay",
		"-1451366515": "DuckDuckGo",
		"-1452159623": "Tecvoz",
		"145291791":   "Bump50:50",
		"1453890729":  "Webmin",
		"-1457536113": "CradlePoint",
		"-1457628171": "Postmark",
		"-1461493576": "tdssuite",
		"1462981117":  "Cyberoam",
		"-1465226128": "Atriuum",
		"-1465479343": "DNN (CMS)",
		"-146638980":  "SharePoint",
		"-1466785234": "Dahua",
		"1466912879":  "CradlePoint Technology (Router)",
		"1467395679":  "Ligowave (network)",
		"-1469020453": "Britecore",
		"-1471633460": "YY",
		"1472757244":  "gov.uk",
		"-1472951746": "MI router",
		"1473157813":  "CyberHound Appliance",
		"1474516533":  "MagicMail Server",
		"-1474875778": "GLPI",
		"-1475308062": "NEBERO Login",
		"1476335317":  "FireEye",
		"-1476120239": "Dreamfactory",
		"-1476384993": "Adobe Connect Central",
		"-1477373218": "Rusonyx",
		"-1477563858": "Arris",
		"-1477694668": "emessage",
		"-147798235":  "Vodafone",
		"-1478010471": "wiki.js",
		"1479202414":  "Arcadyan o2 box (Network)",
		"1479449210":  "UniFi Controller",
		"1482591128":  "Plone",
		"1483097076":  "SyncThru Web Service (Printers)",
		"1485257654":  "SonarQube",
		"1486876794":  "ip3sec",
		"1487164831":  "Metabase",
		"1487209721":  "Kerio Clientless SSL-VPN",
		"-1489194754": "FatPipe",
		"-1489682310": "ICON Time system",
		"1490343308":  "MK-AUTH",
		"-1491070374": "iMM Control Center",
		"1491281975":  "PragmaRX",
		"-1492653156": "Speco technologies",
		"-1492966240": "RADIX",
		"149371702":   "Synology DiskStation",
		"149496700":   "Wisenet",
		"-1495061773": "Metalink",
		"1495420903":  "Arris",
		"1495853641":  "RedHat",
		"1496254733":  "JBoss",
		"1496849154":  "Baidu",
		"-1498185948": "Apple",
		"-1499488123": "Tainy E/HMOD",
		"1500512504":  "leanote",
		"1500747026":  "Powerschool",
		"1501109148":  "Silver Peak",
		"1502215759":  "Mirth",
		"1502482995":  "Indeed",
		"1502815117":  "pgAdmin",
		"1503188865":  "FuelPHP",
		"-1505158120": "Jedox",
		"-1507094812": "MLflow",
		"-1507567067": "Baidu (IP error page)",
		"-1507821392": "Mastodon server",
		"1510327675":  "Zerto",
		"1510423261":  "maxView Storage Manager",
		"151225607":   "Webif^2",
		"-1516177449": "SAP",
		"1517005673":  "Redmine",
		"-1517793453": "CentOS",
		"1522616207":  "CentOS",
		"1523210179":  "Gigaset",
		"1523284537":  "Yahoo!",
		"1526179599":  "Foreman",
		"1526980381":  "Intraweb",
		"1528355650":  "Pinterest",
		"-1528414776": "Rumpus",
		"1529113957":  "Milestone Web Client",
		"1529794015":  "Manaba",
		"1530688358":  "Argon Technologies",
		"-1532282299": "IP Camera",
		"1537743656":  "DropBox",
		"1540037626":  "VKontakte",
		"1540323611":  "Zentyal",
		"-1540609146": "KeystoneJS",
		"1540720428":  "SysAid",
		"1544230796":  "cPanel Login",
		"1544596682":  "TelevisGo",
		"-1544605732": "Amazon",
		"1545970007":  "Magnolia",
		"-1546574541": "Sonatype Nexus Repository Manager",
		"-1547576879": "Saia Burgess Controls - PCD",
		"-1549154783": "Bank of America",
		"1552601947":  "FlexNet",
		"1552860581":  "Elastic (Database)",
		"1553502622":  "A10 Networks",
		"1555083561":  "Linux",
		"-155518807":  "ilsonline",
		"-1561873722": "Nginx",
		"1561951501":  "Zyxel",
		"156312019":   "Technicolor / Thomson Speedtouch (Network / ADSL)",
		"1563218345":  "Postfix Admin",
		"-1564279764": "SuperHub 2",
		"1565357565":  "wallabag",
		"1565485509":  "Amazon AWS",
		"-1566222852": "Nice CXone",
		"-1566499661": "NSFocus",
		"1566659252":  "KBlue",
		"1568516651":  "Acronis",
		"-1569190618": "TransIP",
		"-1571472432": "Sierra Wireless Ace Manager (Airlink)",
		"-157270335":  "Barracuda Firewall",
		"1574384512":  "Storybook",
		"-1576282848": "GoCD",
		"-1577363222": "AMP - Application Management Panel",
		"-1580354544": "MailStore Web Access",
		"-1581907337": "Atlassian - JIRA",
		"15831193":    "WatchGuard",
		"1583275852":  "Spectranet",
		"-1584893473": "docker-swarm-visualizer",
		"1585145626":  "netdata dashboard",
		"-1587809317": "Strapi app",
		"-1588746893": "CommuniGate",
		"-1589842876": "Deluge Web UI",
		"1591050260":  "Aipo",
		"1591770950":  "ValidMail",
		"-1592540846": "i-VU Router",
		"-1593402357": "AT&T",
		"-1593651747": "Blackboard",
		"1594377337":  "Technicolor",
		"1597827316":  "PUSR device",
		"1601194732":  "Sophos Cyberoam (appliance)",
		"1603223646":  "Comcast Business",
		"1604149042":  "Synology",
		"-160425702":  "Medallia",
		"1604363273":  "AeroHive Networks",
		"160541240":   "Transparency Analysis Notification Solution",
		"-160596078":  "RSSBUS",
		"-1606103083": "T-Mobile",
		"160634455":   "O2 easy setup",
		"-1607644090": "Bitnami",
		"-1608064257": "Dell",
		"-1608282972": "VMS Broadcast Server",
		"-1610980956": "Drupal",
		"1611729805":  "Elastic (Database)",
		"-1612496354": "Teltonika",
		"-1614186944": "Kelley Blue Book",
		"-1615532813": "Tautulli",
		"-1616115760": "ownCloud",
		"-1616143106": "AXIS (network cameras)",
		"16202868":    "Universal Devices (UD)",
		"-1620734764": "Velocix Asset Portal",
		"1620794131":  "SocketCluster",
		"1623760745":  "bee-queue/arena",
		"162418931":   "Absorb",
		"1627330242":  "Joomla",
		"-1629133697": "IIS Windows Server",
		"1629518721":  "macOS Server (Apple)",
		"-1630354993": "Proofpoint",
		"1632680057":  "LiveConfig",
		"1632780968":  "Université Toulouse 1 Capitole",
		"1635147929":  "Criteo Corp",
		"-1637180585": "Arlo",
		"-1637198354": "FileMaker",
		"-163806709":  "AlphaNet",
		"-1638258999": "Restreamer",
		"1638963529":  "Observium",
		"163842882":   "Cisco Meraki",
		"-1640178715": "Reddit",
		"1640914316":  "Araknis",
		"-1642012180": "MyOmBox",
		"-1642532491": "Atlassian - Confluence",
		"1642701741":  "Vmware Secure File Transfer",
		"164535647":   "Bipiemme Technology",
		"1648531157":  "InfiNet Wireless | WANFleX (Network)",
		"-1649762785": "webflow",
		"-165021467":  "Ambari",
		"-1654229048": "Vivotek (Camera)",
		"-1655234758": "WSO2",
		"-1656662001": "iServer",
		"-1656695885": "iomega NAS",
		"165976831":   "Vodafone (Technicolor)",
		"-1660707210": "Vinyl",
		"-1661331374": "Dedicated Micros",
		"-166151761":  "Abilis (Network/Automation)",
		"-1661672841": "Fexa",
		"-1661746099": "netDocuments",
		"1662176488":  "PV Webserver",
		"1663064917":  "MyChart",
		"-1664635936": "Telnet",
		"-1666561833": "Wildfly",
		"-1667378049": "Chamilo",
		"1668183286":  "Kibana",
		"1668385882":  "Avigilon",
		"1668745903":  "Piwik",
		"1668832054":  "Adform",
		"1669535914":  "SoftGuard Desktop Security Suite",
		"1673203892":  "Oracle",
		"-167656799":  "Drupal",
		"-167663731":  "Bitwarden",
		"-1677255344": "UBNT Router UI",
		"1678170702":  "Asustor",
		"-1678298769": "Kerio Connect WebMail",
		"-1688698891": "SpamExperts",
		"-1697334194": "Univention Portal",
		"-1702393021": "mofinetwork",
		"-1702769256": "Bosch Security Systems (Camera)",
		"1703788174":  "D-Link (router/network)",
		"-1710631084": "Askey Cable Modem",
		"-1723752240": "Microhard Systems",
		"1726027799":  "IBM Server",
		"1732786188":  "Apache",
		"-1734573358": "TC-Group",
		"1734609466":  "JustHost",
		"1735289686":  "Whatsapp",
		"-1738184811": "cacaoweb",
		"-1738727418": "KeepItSafe Management Console",
		"-1745552996": "Arbor Networks",
		"-1748763891": "INSTAR Full-HD IP-Camera",
		"-175283071":  "Dell",
		"1768726119":  "Outlook Web Application",
		"1770799630":  "bintec elmeg",
		"1772087922":  "ASP.net",
		"-1775553655": "Unified Management Console (Polycom)",
		"-1779611449": "Alienvault",
		"1782271534":  "truVision NVR (interlogix)",
		"1786752597":  "wdCP cloud host management system",
		"-178685903":  "Yasni",
		"-1788112745": "PowerMTA monitoring",
		"1802374283":  "LiquidPixels",
		"-1807411396": "Skype",
		"-1810847295": "Sangfor",
		"-1814887000": "Docker",
		"1821549811":  "(Blank) iSpy",
		"-1822098181": "Checkpoint (Gaia)",
		"-182423204":  "netdata dashboard",
		"-1831547740": "Twitter",
		"-183163807":  "Ace",
		"1835479497":  "Technicolor Gateway",
		"1836828108":  "OpenProject",
		"1838417872":  "Freebox OS",
		"-1840324437": "Microsoft 365",
		"1842351293":  "TP-LINK (Network Device)",
		"-1844173284": "Sohu",
		"1848946384":  "GitHub",
		"-1851491385": "Angular",
		"1862132268":  "Gargoyle Router Management Utility",
		"-1863663974": "Airwatch",
		"1876585825":  "ALIBI NVR",
		"1877797890":  "Eltex (Router)",
		"1895360511":  "VMware Horizon",
		"-1897829998": "D-Link (camera)",
		"1911253822":  "UPC Ceska Republica (Network)",
		"-1911513273": "Dovado Router",
		"1913538826":  "Material Dashboard",
		"1914658187":  "CloudFlare",
		"191654058":   "Wordpress Under Construction Icon",
		"1917028407":  "Vue.js",
		"1922032523":  "NEC WebPro",
		"-1922044295": "Mitel Networks (MiCollab End User Portal)",
		"-1926484046": "Sangfor",
		"-1929912510": "NETASQ - Secure / Stormshield",
		"-1933493443": "Residential Gateway",
		"-1935525788": "SmarterMail",
		"1937209448":  "Docker",
		"1941381095":  "openWRT Luci",
		"1941681276":  "Amazon",
		"-1944119648": "TeamCity",
		"-194439630":  "Avtech IP Surveillance (Camera)",
		"-1950415971": "Joomla",
		"1953726032":  "Metabase",
		"1954835352":  "Vesta Hosting Control Panel",
		"-195508437":  "iPECS",
		"-1961046099": "Dgraph Ratel",
		"1966198264":  "OpenERP (now known as Odoo)",
		"1969970750":  "Gitea",
		"1973665246":  "Entrolink",
		"1975413433":  "Sunny WebBox",
		"1985721423":  "WorldClient for Mdaemon",
		"1991136554":  "Instagram",
		"1991562061":  "Niagara Web Server / Tridium",
		"1993518473":  "cPanel Login",
		"-2006308185": "OTRS (Open Ticket Request System)",
		"2006716043":  "Intelbras SA",
		"-2009722838": "React",
		"2019488876":  "Dahua Storm (IP Camera)",
		"-2031183903": "D-Link (Network)",
		"-2042679530": "Alibaba",
		"2047156994":  "Linksys",
		"-2054889066": "Sentora",
		"2055322029":  "Realtek",
		"-2056503929": "iomega NAS",
		"2058391758":  "Debian/Ubuntu Portal",
		"2059618623":  "HP iLO",
		"-2063036701": "Linksys Smart Wi-Fi",
		"2063428236":  "Sentry",
		"2068154487":  "Digium (Switchvox)",
		"-2069014068": "IBM Security Access manager for Web",
		"-2069844696": "Ruckus Wireless",
		"2071993228":  "Nomadix Access Gateway",
		"2072198544":  "Ferozo Panel",
		"2086228042":  "MobileIron",
		"2091258163":  "Microsoft Lync",
		"2099342476":  "PKP (OpenJournalSystems) Public Knowledge Project",
		"2109473187":  "Huawei - Claro",
		"2113497004":  "WiJungle",
		"-2116540786": "bet365",
		"-2117390767": "Spiceworks (panel)",
		"2119159060":  "GMail",
		"2121539357":  "FireEye",
		"2124459909":  "HFS (HTTP File Server)",
		"-2125083197": "Windows Azure",
		"2127152956":  "MailWizz",
		"2128230701":  "Chainpoint",
		"-2133341160": "WordPress Org",
		"-2138771289": "Technicolor",
		"-2140379067": "RoundCube Webmail",
		"2141724739":  "Juniper Device Manager",
		"-2144363468": "HP Printer / Server",
		"-2145085239": "Tenda Web Master",
		"2146763496":  "Mailcow",
		"-219752612":  "FRITZ!Box",
		"-222497010":  "JoyRun",
		"224536051":   "Shenzhen coship electronics co.",
		"225632504":   "Rocket Chat",
		"-235701012":  "Cnservers LLC",
		"239966418":   "Microsoft Outlook",
		"240136437":   "Seagate Technology (NAS)",
		"240606739":   "FireEye",
		"246145559":   "Parse",
		"251106693":   "GPON Home Gateway",
		"-254193850":  "React",
		"252728887":   "DD WRT (DD-WRT milli_httpd)",
		"255892555":   "wdCP cloud host management system",
		"-256828986":  "iDirect Canada (Network Management)",
		"-266008933":  "SAP Netweaver",
		"-267431135":  "Kibana",
		"-271448102":  "iKuai Networks",
		"-276759139":  "Chef Automate",
		"-277464596":  "AEM Screens",
		"281559989":   "Huawei",
		"283740897":   "Intelbras SA",
		"29056450":    "Geneko",
		"-291579889":  "WS server test page",
		"-297069493":  "Apache Tomcat",
		"-299287097":  "Cisco Router",
		"-299324825":  "Lupus Electronics XT",
		"-305179312":  "Atlassian - Confluence",
		"309020573":   "PayPal",
		"314969666":   "Amazon AWS",
		"-318947884":  "Palo Alto Networks",
		"-318968846":  "ngX-Rocket",
		"31972968":    "Dlink Webcam",
		"321591353":   "Node-RED",
		"321909464":   "Airwatch",
		"322531336":   "iomega NAS",
		"-325082670":  "LANCOM Systems",
		"-329747115":  "C-Lodop",
		"331870709":   "iomega NAS",
		"-332324409":  "STARFACE VoIP Software",
		"-333791179":  "Adobe Campaign Classic",
		"-335153896":  "Traccar GPS tracking",
		"-335242539":  "f5 Big IP",
		"-336242473":  "Siemens OZW772",
		"-342262483":  "Combivox",
		"-35107086":   "WorldClient for Mdaemon",
		"-355305208":  "D-Link (camera)",
		"-359621743":  "Intelbras Wireless",
		"-360566773":  "ARRIS (Network)",
		"362091310":   "MobileIron",
		"363324987":   "Dell SonicWALL",
		"366524387":   "Joomla",
		"-368490461":  "Entronix Energy Management Platform",
		"-373674173":  "Digital Keystone (DK)",
		"-374133142":  "Flower",
		"-374235895":  "Ossia (Provision SR) | Webcam/IP Camera",
		"-375623619":  "bintec elmeg",
		"381100274":   "Moxapass ioLogik Remote Ethernet I/O Server",
		"-38580010":   "Magento",
		"-386189083":  "aaPanel",
		"-38705358":   "Reolink",
		"-393788031":  "Flussonic (Video Streaming)",
		"396533629":   "OpenVPN",
		"-398568076":  "Wikipedia",
		"-401934945":  "iomega NAS",
		"40617830":    "Frontier Communications",
		"420473080":   "Exostar - Managed Access Gateway",
		"-421986013":  "Homegrown Website Hosting",
		"-429287806":  "Ebay",
		"430582574":   "SmartPing",
		"-43161126":   "Slack",
		"432733105":   "Pi Star",
		"-435817905":  "Cambium Networks",
		"-438482901":  "Moodle",
		"-440644339":  "Zyxel ZyWALL",
		"442749392":   "Microsoft OWA",
		"443944613":   "WAMPSERVER",
		"-450254253":  "idera",
		"-459291760":  "Workday",
		"459900502":   "ZTE Corporation (Gateway/Appliance)",
		"462223993":   "Jeedom (home automation)",
		"-466504476":  "Kerio Control Firewall",
		"475379699":   "Axcient Replibit Management Server",
		"476213314":   "Exacq",
		"-476231906":  "phpMyAdmin",
		"479413330":   "Webmin",
		"483383992":   "ISPConfig",
		"-484708885":  "Zyxel ZyWALL",
		"489340156":   "Smart/OS",
		"494866796":   "Aplikasi",
		"-50306417":   "Kyocera (Printer)",
		"-505448917":  "Discuz!",
		"509789953":   "Farming Simulator Dedicated Server",
		"-510925599":  "Windows (Microsoft Corp)",
		"512590457":   "Trendnet IP camera",
		"516963061":   "Gitlab",
		"517158172":   "D-Link (router/network)",
		"-519765377":  "Parallels Plesk Panel",
		"-520888198":  "Blue Iris (Webcam)",
		"-532394952":  "CX",
		"538323054":   "NethServer Enterprise",
		"538585915":   "Lenel",
		"541087742":   "LiquidFiles",
		"545827989":   "MobileIron",
		"-547019147":  "Fedora Server",
		"547025948":   "Grafana",
		"5471989":     "Netcom Technology",
		"547282364":   "Keenetic",
		"547474373":   "TOTOLINK (network)",
		"552592949":   "ASUS AiCloud",
		"552597979":   "Sails",
		"552727997":   "Atlassian - JIRA",
		"5542029":     "NetComWireless (Network)",
		"-560297467":  "DVR (Korean)",
		"56079838":    "Okta",
		"-566516473":  "Meriva Security",
		"-569941107":  "Fireware Watchguard",
		"575613323":   "Canvas LMS (Learning Management)",
		"577446824":   "Bluehost",
		"579239725":   "Metasploit",
		"586998417":   "Nginx",
		"-587741716":  "ADB Broadband (Network)",
		"-590892202":  "Surfilter SSL VPN Portal",
		"593396886":   "StackOverflow",
		"-594256627":  "Netis (network devices)",
		"-600508822":  "iomega NAS",
		"602431586":   "Palo Alto Login Portal",
		"603314":      "Redmine",
		"607846949":   "Hitron Technologies",
		"-609520537":  "OpenGeo Suite",
		"-613216179":  "iomega NAS",
		"-614457039":  "Plesk",
		"-617743584":  "Odoo",
		"-624805968":  "Cloudinary",
		"-625364318":  "OkoFEN Pellematic",
		"628535358":   "Atlassian",
		"-629047854":  "Jetty 404",
		"-630493013":  "DokuWiki",
		"-631002664":  "Kerio Control Firewall",
		"631108382":   "SonicWALL",
		"-632070065":  "Apache Haus",
		"-632583950":  "Shoutcast Server",
		"-644617577":  "SmartLAN/G",
		"648382619":   "Drupal",
		"-649378830":  "WHM",
		"-652508439":  "Parallels Plesk Panel",
		"-655683626":  "PRTG Network Monitor",
		"-656811182":  "Jboss",
		"656868270":   "iomega NAS",
		"661332347":   "MOBOTIX Camera",
		"669847331":   "Unify",
		"671221099":   "innovaphone",
		"-675839242":  "openWRT Luci",
		"-676077969":  "Niagara Web Server",
		"-677167908":  "Kerio Connect (Webmail)",
		"-687783882":  "ClaimTime (Ramsell Public Health & Safety)",
		"-689902428":  "iomega NAS",
		"-692947551":  "Ruijie Networks (Login)",
		"-693082538":  "openmediavault (NAS)",
		"-696586294":  "LinkedIn",
		"693122507":   "WordPress",
		"-697231354":  "Ubiquiti - AirOS",
		"-702384832":  "TCN",
		"705143395":   "Atlassian",
		"706602230":   "VisualSVN Server",
		"708578229":   "Google",
		"711742418":   "Hitron Technologies Inc.",
		"716989053":   "Amazon AWS",
		"72005642":    "RemObjects SDK / Remoting SDK for .NET HTTP Server Microsoft",
		"-723685921":  "Oracle Cloud",
		"726817668":   "KeyHelp (Keyweb AG)",
		"727253975":   "Paradox IP Module",
		"728788645":   "IBM Notes",
		"731374291":   "HFS (HTTP File Server)",
		"-736276076":  "MyASP",
		"-740211187":  "Bing",
		"743365239":   "Atlassian",
		"74935566":    "WindRiver-WebServer",
		"75230260":    "Kibana",
		"758890177":   "Tumblr",
		"-759108386":  "Tongda",
		"-759754862":  "Kibana",
		"76658403":    "TheTradeDesk",
		"-766957661":  "MDaemon Webmail",
		"768231242":   "JAWS Web Server (IP Camera)",
		"768816037":   "UniFi Video Controller (airVision)",
		"77044418":    "Polycom",
		"-771764544":  "Parallels Plesk Panel",
		"774252049":   "FastPanel Hosting",
		"784872924":   "Lucee!",
		"786476039":   "AppsFlyer",
		"786533217":   "OpenStack",
		"788771792":   "Airwatch",
		"794809961":   "CheckPoint",
		"804949239":   "Cisco Meraki Dashboard",
		"812385209":   "Solarwinds Serv-U FTP Server",
		"81586312":    "Jenkins",
		"-816821232":  "GitLab",
		"829321644":   "BOMGAR Support Portal",
		"-831826827":  "NOS Router",
		"833190513":   "Dahua Storm (IP Camera)",
		"-842192932":  "FireEye",
		"855273746":   "JIRA",
		"-85666451":   "DOX - Document Exchange Service",
		"86919334":    "ServiceNow",
		"-873627015":  "HeroSpeed Digital Technology Co. (NVR/IPC/XVR)",
		"878647854":   "BIG-IP",
		"-878891718":  "Twonky Server (Media Streaming)",
		"882208493":   "Lantronix (Spider)",
		"-882760066":  "ZyXEL (Network)",
		"-884776764":  "Huawei (Network)",
		"892542951":   "Zabbix",
		"89321398":    "Askey Cable Modem",
		"-895890586":  "PLEX Server",
		"-895963602":  "Jupyter Notebook",
		"896412703":   "IW",
		"899457975":   "Cisco",
		"90066852":    "JAWS Web Server (IP Camera)",
		"902521196":   "Netflix",
		"903086190":   "Honeywell",
		"904434662":   "Loxone (Automation)",
		"905744673":   "HP Printer / Server",
		"905796143":   "Medallia",
		"90680708":    "Domoticz (Home Automation)",
		"916642917":   "Multilaser",
		"917966895":   "Gogs",
		"920338972":   "Linode",
		"-923088984":  "OpenStack",
		"-923693877":  "motionEye (camera)",
		"926501571":   "Handle Proxy",
		"929825723":   "WAMPSERVER",
		"936297245":   "Twitch",
		"937999361":   "JBoss Application Server 7",
		"938616453":   "Mersive Solstice",
		"943925975":   "ZyXEL",
		"944969688":   "Deluge",
		"945408572":   "Fortinet - Forticlient",
		"95271369":    "FireEye",
		"955369722":   "Sitecore",
		"-956471263":  "Web Client Pro",
		"966563415":   "WordPress Org",
		"967636089":   "MobileIron",
		"970132176":   "3CX Phone System",
		"-972810761":  "HostMonster - Web hosting",
		"97604680":    "Tandberg",
		"-976235259":  "Roundcube Webmail",
		"-978656757":  "NETIASPOT (Network)",
		"979634648":   "StruxureWare (Schneider Electric)",
		"980692677":   "Cake PHP",
		"-981606721":  "Plesk",
		"981867722":   "Atlassian - JIRA",
		"-986678507":  "ISP Manager",
		"-986816620":  "OpenRG",
		"987967490":   "Huawei (Network)",
		"988422585":   "CapRover",
		"-991123252":  "VMware Horizon",
		"99395752":    "Slack",
		"99432374":    "MDaemon Remote Administration",
		"998138196":   "iomega NAS",
		"999357577":   "Hikvision camera",
	}

	ErrHashNotFound    = errors.New("hash not found")
	ErrHashNotMatching = errors.New("hash not matching hash provided")
)

// CheckFavicon checks if faviconHash is present in the database. If hash (slice) is not empty,
// it checks also if that faviconHash is one of the inputted hashes.
// If faviconHash is not found, an error is returned.
func CheckFavicon(faviconHash string, hash goflags.StringSlice, url ...string) (string, error) {
	if k, ok := db[faviconHash]; ok {
		if len(hash) != 0 {
			if contains(hash, faviconHash) {
				return k, nil
			}

			return "", fmt.Errorf("[%s] %s %w", faviconHash, url, ErrHashNotMatching)
		}

		return k, nil
	}

	if len(url) == 0 {
		return "", fmt.Errorf("%w", ErrHashNotFound)
	}

	return "", fmt.Errorf("[%s] %s %w", faviconHash, url, ErrHashNotFound)
}
