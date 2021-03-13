package krakenapi

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"time"
)

// trade pairs constants
const (
	AAVEAUD   = "AAVEAUD"
	AAVEETH   = "AAVEETH"
	AAVEEUR   = "AAVEEUR"
	AAVEGBP   = "AAVEGBP"
	AAVEUSD   = "AAVEUSD"
	AAVEXBT   = "AAVEXBT"
	ADAAUD    = "ADAAUD"
	ADAETH    = "ADAETH"
	ADAEUR    = "ADAEUR"
	ADAGBP    = "ADAGBP"
	ADAUSD    = "ADAUSD"
	ADAUSDT   = "ADAUSDT"
	ADAXBT    = "ADAXBT"
	ALGOETH   = "ALGOETH"
	ALGOEUR   = "ALGOEUR"
	ALGOGBP   = "ALGOGBP"
	ALGOUSD   = "ALGOUSD"
	ALGOXBT   = "ALGOXBT"
	ANTETH    = "ANTETH"
	ANTEUR    = "ANTEUR"
	ANTUSD    = "ANTUSD"
	ANTXBT    = "ANTXBT"
	ATOMAUD   = "ATOMAUD"
	ATOMETH   = "ATOMETH"
	ATOMEUR   = "ATOMEUR"
	ATOMGBP   = "ATOMGBP"
	ATOMUSD   = "ATOMUSD"
	ATOMXBT   = "ATOMXBT"
	AUDJPY    = "AUDJPY"
	AUDUSD    = "AUDUSD"
	BALETH    = "BALETH"
	BALEUR    = "BALEUR"
	BALUSD    = "BALUSD"
	BALXBT    = "BALXBT"
	BATETH    = "BATETH"
	BATEUR    = "BATEUR"
	BATUSD    = "BATUSD"
	BATXBT    = "BATXBT"
	BCHAUD    = "BCHAUD"
	BCHETH    = "BCHETH"
	BCHEUR    = "BCHEUR"
	BCHGBP    = "BCHGBP"
	BCHJPY    = "BCHJPY"
	BCHUSD    = "BCHUSD"
	BCHUSDT   = "BCHUSDT"
	BCHXBT    = "BCHXBT"
	COMPETH   = "COMPETH"
	COMPEUR   = "COMPEUR"
	COMPUSD   = "COMPUSD"
	COMPXBT   = "COMPXBT"
	CRVETH    = "CRVETH"
	CRVEUR    = "CRVEUR"
	CRVUSD    = "CRVUSD"
	CRVXBT    = "CRVXBT"
	DAIEUR    = "DAIEUR"
	DAIUSD    = "DAIUSD"
	DAIUSDT   = "DAIUSDT"
	DASHEUR   = "DASHEUR"
	DASHUSD   = "DASHUSD"
	DASHXBT   = "DASHXBT"
	DOTAUD    = "DOTAUD"
	DOTETH    = "DOTETH"
	DOTEUR    = "DOTEUR"
	DOTGBP    = "DOTGBP"
	DOTUSD    = "DOTUSD"
	DOTUSDT   = "DOTUSDT"
	DOTXBT    = "DOTXBT"
	EOSETH    = "EOSETH"
	EOSEUR    = "EOSEUR"
	EOSUSD    = "EOSUSD"
	EOSUSDT   = "EOSUSDT"
	EOSXBT    = "EOSXBT"
	ETH2SETH  = "ETH2SETH"
	ETHAUD    = "ETHAUD"
	ETHCHF    = "ETHCHF"
	ETHDAI    = "ETHDAI"
	ETHUSDC   = "ETHUSDC"
	ETHUSDT   = "ETHUSDT"
	EURAUD    = "EURAUD"
	EURCAD    = "EURCAD"
	EURCHF    = "EURCHF"
	EURGBP    = "EURGBP"
	EURJPY    = "EURJPY"
	EWTEUR    = "EWTEUR"
	EWTGBP    = "EWTGBP"
	EWTUSD    = "EWTUSD"
	EWTXBT    = "EWTXBT"
	FILAUD    = "FILAUD"
	FILETH    = "FILETH"
	FILEUR    = "FILEUR"
	FILGBP    = "FILGBP"
	FILUSD    = "FILUSD"
	FILXBT    = "FILXBT"
	FLOWETH   = "FLOWETH"
	FLOWEUR   = "FLOWEUR"
	FLOWGBP   = "FLOWGBP"
	FLOWUSD   = "FLOWUSD"
	FLOWXBT   = "FLOWXBT"
	GNOETH    = "GNOETH"
	GNOEUR    = "GNOEUR"
	GNOUSD    = "GNOUSD"
	GNOXBT    = "GNOXBT"
	GRTAUD    = "GRTAUD"
	GRTETH    = "GRTETH"
	GRTEUR    = "GRTEUR"
	GRTGBP    = "GRTGBP"
	GRTUSD    = "GRTUSD"
	GRTXBT    = "GRTXBT"
	ICXETH    = "ICXETH"
	ICXEUR    = "ICXEUR"
	ICXUSD    = "ICXUSD"
	ICXXBT    = "ICXXBT"
	KAVAETH   = "KAVAETH"
	KAVAEUR   = "KAVAEUR"
	KAVAUSD   = "KAVAUSD"
	KAVAXBT   = "KAVAXBT"
	KEEPETH   = "KEEPETH"
	KEEPEUR   = "KEEPEUR"
	KEEPUSD   = "KEEPUSD"
	KEEPXBT   = "KEEPXBT"
	KNCETH    = "KNCETH"
	KNCEUR    = "KNCEUR"
	KNCUSD    = "KNCUSD"
	KNCXBT    = "KNCXBT"
	KSMAUD    = "KSMAUD"
	KSMETH    = "KSMETH"
	KSMEUR    = "KSMEUR"
	KSMGBP    = "KSMGBP"
	KSMUSD    = "KSMUSD"
	KSMXBT    = "KSMXBT"
	LINKAUD   = "LINKAUD"
	LINKETH   = "LINKETH"
	LINKEUR   = "LINKEUR"
	LINKGBP   = "LINKGBP"
	LINKUSD   = "LINKUSD"
	LINKUSDT  = "LINKUSDT"
	LINKXBT   = "LINKXBT"
	LSKETH    = "LSKETH"
	LSKEUR    = "LSKEUR"
	LSKUSD    = "LSKUSD"
	LSKXBT    = "LSKXBT"
	LTCAUD    = "LTCAUD"
	LTCETH    = "LTCETH"
	LTCGBP    = "LTCGBP"
	LTCUSDT   = "LTCUSDT"
	MANAETH   = "MANAETH"
	MANAEUR   = "MANAEUR"
	MANAUSD   = "MANAUSD"
	MANAXBT   = "MANAXBT"
	NANOETH   = "NANOETH"
	NANOEUR   = "NANOEUR"
	NANOUSD   = "NANOUSD"
	NANOXBT   = "NANOXBT"
	OCEANEUR  = "OCEANEUR"
	OCEANGBP  = "OCEANGBP"
	OCEANUSD  = "OCEANUSD"
	OCEANXBT  = "OCEANXBT"
	OMGETH    = "OMGETH"
	OMGEUR    = "OMGEUR"
	OMGUSD    = "OMGUSD"
	OMGXBT    = "OMGXBT"
	OXTETH    = "OXTETH"
	OXTEUR    = "OXTEUR"
	OXTUSD    = "OXTUSD"
	OXTXBT    = "OXTXBT"
	PAXGETH   = "PAXGETH"
	PAXGEUR   = "PAXGEUR"
	PAXGUSD   = "PAXGUSD"
	PAXGXBT   = "PAXGXBT"
	QTUMETH   = "QTUMETH"
	QTUMEUR   = "QTUMEUR"
	QTUMUSD   = "QTUMUSD"
	QTUMXBT   = "QTUMXBT"
	REPV2ETH  = "REPV2ETH"
	REPV2EUR  = "REPV2EUR"
	REPV2USD  = "REPV2USD"
	REPV2XBT  = "REPV2XBT"
	SCETH     = "SCETH"
	SCEUR     = "SCEUR"
	SCUSD     = "SCUSD"
	SCXBT     = "SCXBT"
	SNXAUD    = "SNXAUD"
	SNXETH    = "SNXETH"
	SNXEUR    = "SNXEUR"
	SNXGBP    = "SNXGBP"
	SNXUSD    = "SNXUSD"
	SNXXBT    = "SNXXBT"
	STORJETH  = "STORJETH"
	STORJEUR  = "STORJEUR"
	STORJUSD  = "STORJUSD"
	STORJXBT  = "STORJXBT"
	TBTCETH   = "TBTCETH"
	TBTCEUR   = "TBTCEUR"
	TBTCUSD   = "TBTCUSD"
	TBTCXBT   = "TBTCXBT"
	TRXETH    = "TRXETH"
	TRXEUR    = "TRXEUR"
	TRXUSD    = "TRXUSD"
	TRXXBT    = "TRXXBT"
	UNIETH    = "UNIETH"
	UNIEUR    = "UNIEUR"
	UNIUSD    = "UNIUSD"
	UNIXBT    = "UNIXBT"
	USDCAUD   = "USDCAUD"
	USDCEUR   = "USDCEUR"
	USDCGBP   = "USDCGBP"
	USDCHF    = "USDCHF"
	USDCUSD   = "USDCUSD"
	USDCUSDT  = "USDCUSDT"
	USDTAUD   = "USDTAUD"
	USDTCAD   = "USDTCAD"
	USDTCHF   = "USDTCHF"
	USDTEUR   = "USDTEUR"
	USDTGBP   = "USDTGBP"
	USDTJPY   = "USDTJPY"
	USDTZUSD  = "USDTZUSD"
	WAVESETH  = "WAVESETH"
	WAVESEUR  = "WAVESEUR"
	WAVESUSD  = "WAVESUSD"
	WAVESXBT  = "WAVESXBT"
	XBTAUD    = "XBTAUD"
	XBTCHF    = "XBTCHF"
	XBTDAI    = "XBTDAI"
	XBTUSDC   = "XBTUSDC"
	XBTUSDT   = "XBTUSDT"
	XDGEUR    = "XDGEUR"
	XDGUSD    = "XDGUSD"
	XETCXETH  = "XETCXETH"
	XETCXXBT  = "XETCXXBT"
	XETCZEUR  = "XETCZEUR"
	XETCZUSD  = "XETCZUSD"
	XETHXXBT  = "XETHXXBT"
	XETHXXBTD = "XETHXXBTD"
	XETHZCAD  = "XETHZCAD"
	XETHZCADD = "XETHZCADD"
	XETHZEUR  = "XETHZEUR"
	XETHZEURD = "XETHZEURD"
	XETHZGBP  = "XETHZGBP"
	XETHZGBPD = "XETHZGBPD"
	XETHZJPY  = "XETHZJPY"
	XETHZJPYD = "XETHZJPYD"
	XETHZUSD  = "XETHZUSD"
	XETHZUSDD = "XETHZUSDD"
	XLTCXXBT  = "XLTCXXBT"
	XLTCZEUR  = "XLTCZEUR"
	XLTCZJPY  = "XLTCZJPY"
	XLTCZUSD  = "XLTCZUSD"
	XMLNXETH  = "XMLNXETH"
	XMLNXXBT  = "XMLNXXBT"
	XMLNZEUR  = "XMLNZEUR"
	XMLNZUSD  = "XMLNZUSD"
	XREPXETH  = "XREPXETH"
	XREPXXBT  = "XREPXXBT"
	XREPZEUR  = "XREPZEUR"
	XREPZUSD  = "XREPZUSD"
	XRPAUD    = "XRPAUD"
	XRPETH    = "XRPETH"
	XRPGBP    = "XRPGBP"
	XRPUSDT   = "XRPUSDT"
	XTZAUD    = "XTZAUD"
	XTZETH    = "XTZETH"
	XTZEUR    = "XTZEUR"
	XTZGBP    = "XTZGBP"
	XTZUSD    = "XTZUSD"
	XTZXBT    = "XTZXBT"
	XXBTZCAD  = "XXBTZCAD"
	XXBTZCADD = "XXBTZCADD"
	XXBTZEUR  = "XXBTZEUR"
	XXBTZEURD = "XXBTZEURD"
	XXBTZGBP  = "XXBTZGBP"
	XXBTZGBPD = "XXBTZGBPD"
	XXBTZJPY  = "XXBTZJPY"
	XXBTZJPYD = "XXBTZJPYD"
	XXBTZUSD  = "XXBTZUSD"
	XXBTZUSDD = "XXBTZUSDD"
	XXDGXXBT  = "XXDGXXBT"
	XXLMXXBT  = "XXLMXXBT"
	XXLMZAUD  = "XXLMZAUD"
	XXLMZEUR  = "XXLMZEUR"
	XXLMZGBP  = "XXLMZGBP"
	XXLMZUSD  = "XXLMZUSD"
	XXMRXXBT  = "XXMRXXBT"
	XXMRZEUR  = "XXMRZEUR"
	XXMRZUSD  = "XXMRZUSD"
	XXRPXXBT  = "XXRPXXBT"
	XXRPZCAD  = "XXRPZCAD"
	XXRPZEUR  = "XXRPZEUR"
	XXRPZJPY  = "XXRPZJPY"
	XXRPZUSD  = "XXRPZUSD"
	XZECXXBT  = "XZECXXBT"
	XZECZEUR  = "XZECZEUR"
	XZECZUSD  = "XZECZUSD"
	YFIAUD    = "YFIAUD"
	YFIETH    = "YFIETH"
	YFIEUR    = "YFIEUR"
	YFIGBP    = "YFIGBP"
	YFIUSD    = "YFIUSD"
	YFIXBT    = "YFIXBT"
	ZEURZUSD  = "ZEURZUSD"
	ZGBPZUSD  = "ZGBPZUSD"
	ZUSDZCAD  = "ZUSDZCAD"
	ZUSDZJPY  = "ZUSDZJPY"
)

// actions constants
const (
	BUY    = "b"
	SELL   = "s"
	MARKET = "m"
	LIMIT  = "l"
)

// KrakenResponse wraps the Kraken API JSON response
type KrakenResponse struct {
	Error  []string    `json:"error"`
	Result interface{} `json:"result"`
}

// TimeResponse represents the server's time
type TimeResponse struct {
	// Unix timestamp
	Unixtime int64
	// RFC 1123 time format
	Rfc1123 string
}

// AssetPairsResponse includes asset pair information
type AssetPairsResponse struct {
	AAVEAUD   AssetPairInfo
	AAVEETH   AssetPairInfo
	AAVEEUR   AssetPairInfo
	AAVEGBP   AssetPairInfo
	AAVEUSD   AssetPairInfo
	AAVEXBT   AssetPairInfo
	ADAAUD    AssetPairInfo
	ADAETH    AssetPairInfo
	ADAEUR    AssetPairInfo
	ADAGBP    AssetPairInfo
	ADAUSD    AssetPairInfo
	ADAUSDT   AssetPairInfo
	ADAXBT    AssetPairInfo
	ALGOETH   AssetPairInfo
	ALGOEUR   AssetPairInfo
	ALGOGBP   AssetPairInfo
	ALGOUSD   AssetPairInfo
	ALGOXBT   AssetPairInfo
	ANTETH    AssetPairInfo
	ANTEUR    AssetPairInfo
	ANTUSD    AssetPairInfo
	ANTXBT    AssetPairInfo
	ATOMAUD   AssetPairInfo
	ATOMETH   AssetPairInfo
	ATOMEUR   AssetPairInfo
	ATOMGBP   AssetPairInfo
	ATOMUSD   AssetPairInfo
	ATOMXBT   AssetPairInfo
	AUDJPY    AssetPairInfo
	AUDUSD    AssetPairInfo
	BALETH    AssetPairInfo
	BALEUR    AssetPairInfo
	BALUSD    AssetPairInfo
	BALXBT    AssetPairInfo
	BATETH    AssetPairInfo
	BATEUR    AssetPairInfo
	BATUSD    AssetPairInfo
	BATXBT    AssetPairInfo
	BCHAUD    AssetPairInfo
	BCHETH    AssetPairInfo
	BCHEUR    AssetPairInfo
	BCHGBP    AssetPairInfo
	BCHJPY    AssetPairInfo
	BCHUSD    AssetPairInfo
	BCHUSDT   AssetPairInfo
	BCHXBT    AssetPairInfo
	COMPETH   AssetPairInfo
	COMPEUR   AssetPairInfo
	COMPUSD   AssetPairInfo
	COMPXBT   AssetPairInfo
	CRVETH    AssetPairInfo
	CRVEUR    AssetPairInfo
	CRVUSD    AssetPairInfo
	CRVXBT    AssetPairInfo
	DAIEUR    AssetPairInfo
	DAIUSD    AssetPairInfo
	DAIUSDT   AssetPairInfo
	DASHEUR   AssetPairInfo
	DASHUSD   AssetPairInfo
	DASHXBT   AssetPairInfo
	DOTAUD    AssetPairInfo
	DOTETH    AssetPairInfo
	DOTEUR    AssetPairInfo
	DOTGBP    AssetPairInfo
	DOTUSD    AssetPairInfo
	DOTUSDT   AssetPairInfo
	DOTXBT    AssetPairInfo
	EOSETH    AssetPairInfo
	EOSEUR    AssetPairInfo
	EOSUSD    AssetPairInfo
	EOSUSDT   AssetPairInfo
	EOSXBT    AssetPairInfo
	ETH2SETH  AssetPairInfo
	ETHAUD    AssetPairInfo
	ETHCHF    AssetPairInfo
	ETHDAI    AssetPairInfo
	ETHUSDC   AssetPairInfo
	ETHUSDT   AssetPairInfo
	EURAUD    AssetPairInfo
	EURCAD    AssetPairInfo
	EURCHF    AssetPairInfo
	EURGBP    AssetPairInfo
	EURJPY    AssetPairInfo
	EWTEUR    AssetPairInfo
	EWTGBP    AssetPairInfo
	EWTUSD    AssetPairInfo
	EWTXBT    AssetPairInfo
	FILAUD    AssetPairInfo
	FILETH    AssetPairInfo
	FILEUR    AssetPairInfo
	FILGBP    AssetPairInfo
	FILUSD    AssetPairInfo
	FILXBT    AssetPairInfo
	FLOWETH   AssetPairInfo
	FLOWEUR   AssetPairInfo
	FLOWGBP   AssetPairInfo
	FLOWUSD   AssetPairInfo
	FLOWXBT   AssetPairInfo
	GNOETH    AssetPairInfo
	GNOEUR    AssetPairInfo
	GNOUSD    AssetPairInfo
	GNOXBT    AssetPairInfo
	GRTAUD    AssetPairInfo
	GRTETH    AssetPairInfo
	GRTEUR    AssetPairInfo
	GRTGBP    AssetPairInfo
	GRTUSD    AssetPairInfo
	GRTXBT    AssetPairInfo
	ICXETH    AssetPairInfo
	ICXEUR    AssetPairInfo
	ICXUSD    AssetPairInfo
	ICXXBT    AssetPairInfo
	KAVAETH   AssetPairInfo
	KAVAEUR   AssetPairInfo
	KAVAUSD   AssetPairInfo
	KAVAXBT   AssetPairInfo
	KEEPETH   AssetPairInfo
	KEEPEUR   AssetPairInfo
	KEEPUSD   AssetPairInfo
	KEEPXBT   AssetPairInfo
	KNCETH    AssetPairInfo
	KNCEUR    AssetPairInfo
	KNCUSD    AssetPairInfo
	KNCXBT    AssetPairInfo
	KSMAUD    AssetPairInfo
	KSMETH    AssetPairInfo
	KSMEUR    AssetPairInfo
	KSMGBP    AssetPairInfo
	KSMUSD    AssetPairInfo
	KSMXBT    AssetPairInfo
	LINKAUD   AssetPairInfo
	LINKETH   AssetPairInfo
	LINKEUR   AssetPairInfo
	LINKGBP   AssetPairInfo
	LINKUSD   AssetPairInfo
	LINKUSDT  AssetPairInfo
	LINKXBT   AssetPairInfo
	LSKETH    AssetPairInfo
	LSKEUR    AssetPairInfo
	LSKUSD    AssetPairInfo
	LSKXBT    AssetPairInfo
	LTCAUD    AssetPairInfo
	LTCETH    AssetPairInfo
	LTCGBP    AssetPairInfo
	LTCUSDT   AssetPairInfo
	MANAETH   AssetPairInfo
	MANAEUR   AssetPairInfo
	MANAUSD   AssetPairInfo
	MANAXBT   AssetPairInfo
	NANOETH   AssetPairInfo
	NANOEUR   AssetPairInfo
	NANOUSD   AssetPairInfo
	NANOXBT   AssetPairInfo
	OCEANEUR  AssetPairInfo
	OCEANGBP  AssetPairInfo
	OCEANUSD  AssetPairInfo
	OCEANXBT  AssetPairInfo
	OMGETH    AssetPairInfo
	OMGEUR    AssetPairInfo
	OMGUSD    AssetPairInfo
	OMGXBT    AssetPairInfo
	OXTETH    AssetPairInfo
	OXTEUR    AssetPairInfo
	OXTUSD    AssetPairInfo
	OXTXBT    AssetPairInfo
	PAXGETH   AssetPairInfo
	PAXGEUR   AssetPairInfo
	PAXGUSD   AssetPairInfo
	PAXGXBT   AssetPairInfo
	QTUMETH   AssetPairInfo
	QTUMEUR   AssetPairInfo
	QTUMUSD   AssetPairInfo
	QTUMXBT   AssetPairInfo
	REPV2ETH  AssetPairInfo
	REPV2EUR  AssetPairInfo
	REPV2USD  AssetPairInfo
	REPV2XBT  AssetPairInfo
	SCETH     AssetPairInfo
	SCEUR     AssetPairInfo
	SCUSD     AssetPairInfo
	SCXBT     AssetPairInfo
	SNXAUD    AssetPairInfo
	SNXETH    AssetPairInfo
	SNXEUR    AssetPairInfo
	SNXGBP    AssetPairInfo
	SNXUSD    AssetPairInfo
	SNXXBT    AssetPairInfo
	STORJETH  AssetPairInfo
	STORJEUR  AssetPairInfo
	STORJUSD  AssetPairInfo
	STORJXBT  AssetPairInfo
	TBTCETH   AssetPairInfo
	TBTCEUR   AssetPairInfo
	TBTCUSD   AssetPairInfo
	TBTCXBT   AssetPairInfo
	TRXETH    AssetPairInfo
	TRXEUR    AssetPairInfo
	TRXUSD    AssetPairInfo
	TRXXBT    AssetPairInfo
	UNIETH    AssetPairInfo
	UNIEUR    AssetPairInfo
	UNIUSD    AssetPairInfo
	UNIXBT    AssetPairInfo
	USDCAUD   AssetPairInfo
	USDCEUR   AssetPairInfo
	USDCGBP   AssetPairInfo
	USDCHF    AssetPairInfo
	USDCUSD   AssetPairInfo
	USDCUSDT  AssetPairInfo
	USDTAUD   AssetPairInfo
	USDTCAD   AssetPairInfo
	USDTCHF   AssetPairInfo
	USDTEUR   AssetPairInfo
	USDTGBP   AssetPairInfo
	USDTJPY   AssetPairInfo
	USDTZUSD  AssetPairInfo
	WAVESETH  AssetPairInfo
	WAVESEUR  AssetPairInfo
	WAVESUSD  AssetPairInfo
	WAVESXBT  AssetPairInfo
	XBTAUD    AssetPairInfo
	XBTCHF    AssetPairInfo
	XBTDAI    AssetPairInfo
	XBTUSDC   AssetPairInfo
	XBTUSDT   AssetPairInfo
	XDGEUR    AssetPairInfo
	XDGUSD    AssetPairInfo
	XETCXETH  AssetPairInfo
	XETCXXBT  AssetPairInfo
	XETCZEUR  AssetPairInfo
	XETCZUSD  AssetPairInfo
	XETHXXBT  AssetPairInfo
	XETHXXBTD AssetPairInfo
	XETHZCAD  AssetPairInfo
	XETHZCADD AssetPairInfo
	XETHZEUR  AssetPairInfo
	XETHZEURD AssetPairInfo
	XETHZGBP  AssetPairInfo
	XETHZGBPD AssetPairInfo
	XETHZJPY  AssetPairInfo
	XETHZJPYD AssetPairInfo
	XETHZUSD  AssetPairInfo
	XETHZUSDD AssetPairInfo
	XLTCXXBT  AssetPairInfo
	XLTCZEUR  AssetPairInfo
	XLTCZJPY  AssetPairInfo
	XLTCZUSD  AssetPairInfo
	XMLNXETH  AssetPairInfo
	XMLNXXBT  AssetPairInfo
	XMLNZEUR  AssetPairInfo
	XMLNZUSD  AssetPairInfo
	XREPXETH  AssetPairInfo
	XREPXXBT  AssetPairInfo
	XREPZEUR  AssetPairInfo
	XREPZUSD  AssetPairInfo
	XRPAUD    AssetPairInfo
	XRPETH    AssetPairInfo
	XRPGBP    AssetPairInfo
	XRPUSDT   AssetPairInfo
	XTZAUD    AssetPairInfo
	XTZETH    AssetPairInfo
	XTZEUR    AssetPairInfo
	XTZGBP    AssetPairInfo
	XTZUSD    AssetPairInfo
	XTZXBT    AssetPairInfo
	XXBTZCAD  AssetPairInfo
	XXBTZCADD AssetPairInfo
	XXBTZEUR  AssetPairInfo
	XXBTZEURD AssetPairInfo
	XXBTZGBP  AssetPairInfo
	XXBTZGBPD AssetPairInfo
	XXBTZJPY  AssetPairInfo
	XXBTZJPYD AssetPairInfo
	XXBTZUSD  AssetPairInfo
	XXBTZUSDD AssetPairInfo
	XXDGXXBT  AssetPairInfo
	XXLMXXBT  AssetPairInfo
	XXLMZAUD  AssetPairInfo
	XXLMZEUR  AssetPairInfo
	XXLMZGBP  AssetPairInfo
	XXLMZUSD  AssetPairInfo
	XXMRXXBT  AssetPairInfo
	XXMRZEUR  AssetPairInfo
	XXMRZUSD  AssetPairInfo
	XXRPXXBT  AssetPairInfo
	XXRPZCAD  AssetPairInfo
	XXRPZEUR  AssetPairInfo
	XXRPZJPY  AssetPairInfo
	XXRPZUSD  AssetPairInfo
	XZECXXBT  AssetPairInfo
	XZECZEUR  AssetPairInfo
	XZECZUSD  AssetPairInfo
	YFIAUD    AssetPairInfo
	YFIETH    AssetPairInfo
	YFIEUR    AssetPairInfo
	YFIGBP    AssetPairInfo
	YFIUSD    AssetPairInfo
	YFIXBT    AssetPairInfo
	ZEURZUSD  AssetPairInfo
	ZGBPZUSD  AssetPairInfo
	ZUSDZCAD  AssetPairInfo
	ZUSDZJPY  AssetPairInfo
}

// AssetPairInfo represents asset pair information
type AssetPairInfo struct {
	// Alternate pair name
	Altname string `json:"altname"`
	//WebSocket pair name (if available)
	WebSocketPairName string `json:"wsname"`
	// Asset class of base component
	AssetClassBase string `json:"aclass_base"`
	// Asset id of base component
	Base string `json:"base"`
	// Asset class of quote component
	AssetClassQuote string `json:"aclass_quote"`
	// Asset id of quote component
	Quote string `json:"quote"`
	// Volume lot size
	Lot string `json:"lot"`
	// Scaling decimal places for pair
	PairDecimals int `json:"pair_decimals"`
	// Scaling decimal places for volume
	LotDecimals int `json:"lot_decimals"`
	// Amount to multiply lot volume by to get currency volume
	LotMultiplier int `json:"lot_multiplier"`
	// Array of leverage amounts available when buying
	LeverageBuy []float64 `json:"leverage_buy"`
	// Array of leverage amounts available when selling
	LeverageSell []float64 `json:"leverage_sell"`
	// Fee schedule array in [volume, percent fee] tuples
	Fees [][]float64 `json:"fees"`
	// // Maker fee schedule array in [volume, percent fee] tuples (if on maker/taker)
	FeesMaker [][]float64 `json:"fees_maker"`
	// // Volume discount currency
	FeeVolumeCurrency string `json:"fee_volume_currency"`
	// Margin call level
	MarginCall int `json:"margin_call"`
	// Stop-out/Liquidation margin level
	MarginStop int `json:"margin_stop"`
	// Order minimum
	OrderMin float64 `json:"ordermin,string"`
}

// AssetsResponse includes asset informations
type AssetsResponse struct {
	ADA  AssetInfo
	AAVE AssetInfo
	BCH  AssetInfo
	DASH AssetInfo
	EOS  AssetInfo
	GNO  AssetInfo
	KFEE AssetInfo
	QTUM AssetInfo
	USDT AssetInfo
	XDAO AssetInfo
	XETC AssetInfo
	XETH AssetInfo
	XICN AssetInfo
	XLTC AssetInfo
	XMLN AssetInfo
	XNMC AssetInfo
	XREP AssetInfo
	XXBT AssetInfo
	XXDG AssetInfo
	XXLM AssetInfo
	XXMR AssetInfo
	XXRP AssetInfo
	XTZ  AssetInfo
	XXVN AssetInfo
	XZEC AssetInfo
	ZCAD AssetInfo
	ZEUR AssetInfo
	ZGBP AssetInfo
	ZJPY AssetInfo
	ZKRW AssetInfo
	ZUSD AssetInfo
}

// AssetInfo represents an asset information
type AssetInfo struct {
	// Alternate name
	Altname string
	// Asset class
	AssetClass string `json:"aclass"`
	// Scaling decimal places for record keeping
	Decimals int
	// Scaling decimal places for output display
	DisplayDecimals int `json:"display_decimals"`
}

// BalanceResponse represents the account's balances (list of currencies)
type BalanceResponse struct {
	AAVE          float64 `json:"AAVE,string,omitempty"`
	ADA           float64 `json:"ADA,string,omitempty"`
	ALGO          float64 `json:"ALGO,string,omitempty"`
	ANT           float64 `json:"ANT,string,omitempty"`
	ATOM          float64 `json:"ATOM,string,omitempty"`
	ATOM_STAKING  float64 `json:"ATOM.S,string,omitempty"`
	BAL           float64 `json:"BAL,string,omitempty"`
	BAT           float64 `json:"BAT,string,omitempty"`
	BCH           float64 `json:"BCH,string,omitempty"`
	CHF           float64 `json:"CHF,string,omitempty"`
	COMP          float64 `json:"COMP,string,omitempty"`
	CRV           float64 `json:"CRV,string,omitempty"`
	DAI           float64 `json:"DAI,string,omitempty"`
	DASH          float64 `json:"DASH,string,omitempty"`
	DOT           float64 `json:"DOT,string,omitempty"`
	DOT_STAKING   float64 `json:"DOT.S,string,omitempty"`
	EOS           float64 `json:"EOS,string,omitempty"`
	ETH2          float64 `json:"ETH2,string,omitempty"`
	ETH2_STAKING  float64 `json:"ETH2.S,string,omitempty"`
	EUR_HOLD      float64 `json:"EUR.HOLD,string,omitempty"`
	EUR_M         float64 `json:"EUR.M,string,omitempty"`
	EWT           float64 `json:"EWT,string,omitempty"`
	FIL           float64 `json:"FIL,string,omitempty"`
	FLOW          float64 `json:"FLOW,string,omitempty"`
	FLOW_STAKING  float64 `json:"FLOW.S,string,omitempty"`
	FLOWH         float64 `json:"FLOWH,string,omitempty"`
	FLOWH_STAKING float64 `json:"FLOWH.S,string,omitempty"`
	GNO           float64 `json:"GNO,string,omitempty"`
	GRT           float64 `json:"GRT,string,omitempty"`
	ICX           float64 `json:"ICX,string,omitempty"`
	KAVA          float64 `json:"KAVA,string,omitempty"`
	KAVA_STAKING  float64 `json:"KAVA.S,string,omitempty"`
	KEEP          float64 `json:"KEEP,string,omitempty"`
	FEE           float64 `json:"KFEE,string,omitempty"`
	KNC           float64 `json:"KNC,string,omitempty"`
	KSM           float64 `json:"KSM,string,omitempty"`
	KSM_STAKING   float64 `json:"KSM.S,string,omitempty"`
	LINK          float64 `json:"LINK,string,omitempty"`
	LSK           float64 `json:"LSK,string,omitempty"`
	MANA          float64 `json:"MANA,string,omitempty"`
	NANO          float64 `json:"NANO,string,omitempty"`
	OCEAN         float64 `json:"OCEAN,string,omitempty"`
	OMG           float64 `json:"OMG,string,omitempty"`
	OXT           float64 `json:"OXT,string,omitempty"`
	PAXG          float64 `json:"PAXG,string,omitempty"`
	QTUM          float64 `json:"QTUM,string,omitempty"`
	REPV2         float64 `json:"REPV2,string,omitempty"`
	SC            float64 `json:"SC,string,omitempty"`
	SNX           float64 `json:"SNX,string,omitempty"`
	STORJ         float64 `json:"STORJ,string,omitempty"`
	TBTC          float64 `json:"TBTC,string,omitempty"`
	TRX           float64 `json:"TRX,string,omitempty"`
	UNI           float64 `json:"UNI,string,omitempty"`
	USD_HOLD      float64 `json:"USD.HOLD,string,omitempty"`
	USD_M         float64 `json:"USD.M,string,omitempty"`
	USDC          float64 `json:"USDC,string,omitempty"`
	USDT          float64 `json:"USDT,string,omitempty"`
	WAVES         float64 `json:"WAVES,string,omitempty"`
	XBT_M         float64 `json:"XBT.M,string,omitempty"`
	ETC           float64 `json:"XETC,string,omitempty"`
	ETH           float64 `json:"XETH,string,omitempty"`
	LTC           float64 `json:"XLTC,string,omitempty"`
	MLN           float64 `json:"XMLN,string,omitempty"`
	REP           float64 `json:"XREP,string,omitempty"`
	XTZ           float64 `json:"XTZ,string,omitempty"`
	XTZ_STAKING   float64 `json:"XTZ.S,string,omitempty"`
	XBT           float64 `json:"XXBT,string,omitempty"`
	XDG           float64 `json:"XXDG,string,omitempty"`
	XLM           float64 `json:"XXLM,string,omitempty"`
	XMR           float64 `json:"XXMR,string,omitempty"`
	XRP           float64 `json:"XXRP,string,omitempty"`
	ZEC           float64 `json:"XZEC,string,omitempty"`
	YFI           float64 `json:"YFI,string,omitempty"`
	AUD           float64 `json:"ZAUD,string,omitempty"`
	CAD           float64 `json:"ZCAD,string,omitempty"`
	EUR           float64 `json:"ZEUR,string,omitempty"`
	GBP           float64 `json:"ZGBP,string,omitempty"`
	JPY           float64 `json:"ZJPY,string,omitempty"`
	USD           float64 `json:"ZUSD,string,omitempty"`
}

// TradeBalanceResponse struct used as the response for the TradeBalance method
type TradeBalanceResponse struct {
	EquivalentBalance         float64 `json:"eb,string"`
	TradeBalance              float64 `json:"tb,string"`
	MarginOP                  float64 `json:"m,string"`
	UnrealizedNetProfitLossOP float64 `json:"n,string"`
	CostBasisOP               float64 `json:"c,string"`
	CurrentValuationOP        float64 `json:"v,string"`
	Equity                    float64 `json:"e,string"`
	FreeMargin                float64 `json:"mf,string"`
	MarginLevel               float64 `json:"ml,string"`
}

// Fees includes fees information for different currencies
type Fees struct {
	ADACAD   FeeInfo
	ADAETH   FeeInfo
	ADAEUR   FeeInfo
	ADAUSD   FeeInfo
	ADAXBT   FeeInfo
	AAVEUSD  FeeInfo
	BCHEUR   FeeInfo
	BCHUSD   FeeInfo
	BCHXBT   FeeInfo
	DASHEUR  FeeInfo
	DASHUSD  FeeInfo
	DASHXBT  FeeInfo
	EOSETH   FeeInfo
	EOSEUR   FeeInfo
	EOSUSD   FeeInfo
	EOSXBT   FeeInfo
	GNOETH   FeeInfo
	GNOEUR   FeeInfo
	GNOUSD   FeeInfo
	GNOXBT   FeeInfo
	QTUMCAD  FeeInfo
	QTUMETH  FeeInfo
	QTUMEUR  FeeInfo
	QTUMUSD  FeeInfo
	QTUMXBT  FeeInfo
	USDTZUSD FeeInfo
	XETCXETH FeeInfo
	XETCXXBT FeeInfo
	XETCZEUR FeeInfo
	XETCZUSD FeeInfo
	XETHXXBT FeeInfo
	XETHZCAD FeeInfo
	XETHZEUR FeeInfo
	XETHZGBP FeeInfo
	XETHZJPY FeeInfo
	XETHZUSD FeeInfo
	XICNXETH FeeInfo
	XICNXXBT FeeInfo
	XLTCXXBT FeeInfo
	XLTCZEUR FeeInfo
	XLTCZUSD FeeInfo
	XMLNXETH FeeInfo
	XMLNXXBT FeeInfo
	XREPXETH FeeInfo
	XREPXXBT FeeInfo
	XREPZEUR FeeInfo
	XREPZUSD FeeInfo
	XTZCAD   FeeInfo
	XTZETH   FeeInfo
	XTZEUR   FeeInfo
	XTZUSD   FeeInfo
	XTZXBT   FeeInfo
	XXBTZCAD FeeInfo
	XXBTZEUR FeeInfo
	XXBTZGBP FeeInfo
	XXBTZJPY FeeInfo
	XXBTZUSD FeeInfo
	XXDGXXBT FeeInfo
	XXLMXXBT FeeInfo
	XXLMZEUR FeeInfo
	XXLMZUSD FeeInfo
	XXMRXXBT FeeInfo
	XXMRZEUR FeeInfo
	XXMRZUSD FeeInfo
	XXRPXXBT FeeInfo
	XXRPZCAD FeeInfo
	XXRPZEUR FeeInfo
	XXRPZJPY FeeInfo
	XXRPZUSD FeeInfo
	XZECXXBT FeeInfo
	XZECZEUR FeeInfo
	XZECZUSD FeeInfo
}

// FeeInfo represents a fee information
type FeeInfo struct {
	Fee        float64 `json:"fee,string"`
	MinFee     float64 `json:"minfee,string"`
	MaxFee     float64 `json:"maxfee,string"`
	NextFee    float64 `json:"nextfee,string"`
	NextVolume float64 `json:"nextvolume,string"`
	TierVolume float64 `json:"tiervolume,string"`
}

// TradeVolumeResponse represents the response for trade volume
type TradeVolumeResponse struct {
	Volume    float64 `json:"volume,string"`
	Currency  string  `json:"currency"`
	Fees      Fees    `json:"fees"`
	FeesMaker Fees    `json:"fees_maker"`
}

// TickerResponse includes the requested ticker pairs
type TickerResponse struct {
	ADACAD    PairTickerInfo
	AAVEAUD   PairTickerInfo
	AAVEETH   PairTickerInfo
	AAVEEUR   PairTickerInfo
	AAVEGBP   PairTickerInfo
	AAVEUSD   PairTickerInfo
	AAVEXBT   PairTickerInfo
	ADAAUD    PairTickerInfo
	ADAETH    PairTickerInfo
	ADAEUR    PairTickerInfo
	ADAGBP    PairTickerInfo
	ADAUSD    PairTickerInfo
	ADAUSDT   PairTickerInfo
	ADAXBT    PairTickerInfo
	ALGOETH   PairTickerInfo
	ALGOEUR   PairTickerInfo
	ALGOGBP   PairTickerInfo
	ALGOUSD   PairTickerInfo
	ALGOXBT   PairTickerInfo
	ANTETH    PairTickerInfo
	ANTEUR    PairTickerInfo
	ANTUSD    PairTickerInfo
	ANTXBT    PairTickerInfo
	ATOMAUD   PairTickerInfo
	ATOMETH   PairTickerInfo
	ATOMEUR   PairTickerInfo
	ATOMGBP   PairTickerInfo
	ATOMUSD   PairTickerInfo
	ATOMXBT   PairTickerInfo
	AUDJPY    PairTickerInfo
	AUDUSD    PairTickerInfo
	BALETH    PairTickerInfo
	BALEUR    PairTickerInfo
	BALUSD    PairTickerInfo
	BALXBT    PairTickerInfo
	BATETH    PairTickerInfo
	BATEUR    PairTickerInfo
	BATUSD    PairTickerInfo
	BATXBT    PairTickerInfo
	BCHAUD    PairTickerInfo
	BCHETH    PairTickerInfo
	BCHEUR    PairTickerInfo
	BCHGBP    PairTickerInfo
	BCHJPY    PairTickerInfo
	BCHUSD    PairTickerInfo
	BCHUSDT   PairTickerInfo
	BCHXBT    PairTickerInfo
	COMPETH   PairTickerInfo
	COMPEUR   PairTickerInfo
	COMPUSD   PairTickerInfo
	COMPXBT   PairTickerInfo
	CRVETH    PairTickerInfo
	CRVEUR    PairTickerInfo
	CRVUSD    PairTickerInfo
	CRVXBT    PairTickerInfo
	DAIEUR    PairTickerInfo
	DAIUSD    PairTickerInfo
	DAIUSDT   PairTickerInfo
	DASHEUR   PairTickerInfo
	DASHUSD   PairTickerInfo
	DASHXBT   PairTickerInfo
	DOTAUD    PairTickerInfo
	DOTETH    PairTickerInfo
	DOTEUR    PairTickerInfo
	DOTGBP    PairTickerInfo
	DOTUSD    PairTickerInfo
	DOTUSDT   PairTickerInfo
	DOTXBT    PairTickerInfo
	EOSETH    PairTickerInfo
	EOSEUR    PairTickerInfo
	EOSUSD    PairTickerInfo
	EOSUSDT   PairTickerInfo
	EOSXBT    PairTickerInfo
	ETH2SETH  PairTickerInfo
	ETHAUD    PairTickerInfo
	ETHCHF    PairTickerInfo
	ETHDAI    PairTickerInfo
	ETHUSDC   PairTickerInfo
	ETHUSDT   PairTickerInfo
	EURAUD    PairTickerInfo
	EURCAD    PairTickerInfo
	EURCHF    PairTickerInfo
	EURGBP    PairTickerInfo
	EURJPY    PairTickerInfo
	EWTEUR    PairTickerInfo
	EWTGBP    PairTickerInfo
	EWTUSD    PairTickerInfo
	EWTXBT    PairTickerInfo
	FILAUD    PairTickerInfo
	FILETH    PairTickerInfo
	FILEUR    PairTickerInfo
	FILGBP    PairTickerInfo
	FILUSD    PairTickerInfo
	FILXBT    PairTickerInfo
	FLOWETH   PairTickerInfo
	FLOWEUR   PairTickerInfo
	FLOWGBP   PairTickerInfo
	FLOWUSD   PairTickerInfo
	FLOWXBT   PairTickerInfo
	GNOETH    PairTickerInfo
	GNOEUR    PairTickerInfo
	GNOUSD    PairTickerInfo
	GNOXBT    PairTickerInfo
	GRTAUD    PairTickerInfo
	GRTETH    PairTickerInfo
	GRTEUR    PairTickerInfo
	GRTGBP    PairTickerInfo
	GRTUSD    PairTickerInfo
	GRTXBT    PairTickerInfo
	ICXETH    PairTickerInfo
	ICXEUR    PairTickerInfo
	ICXUSD    PairTickerInfo
	ICXXBT    PairTickerInfo
	KAVAETH   PairTickerInfo
	KAVAEUR   PairTickerInfo
	KAVAUSD   PairTickerInfo
	KAVAXBT   PairTickerInfo
	KEEPETH   PairTickerInfo
	KEEPEUR   PairTickerInfo
	KEEPUSD   PairTickerInfo
	KEEPXBT   PairTickerInfo
	KNCETH    PairTickerInfo
	KNCEUR    PairTickerInfo
	KNCUSD    PairTickerInfo
	KNCXBT    PairTickerInfo
	KSMAUD    PairTickerInfo
	KSMETH    PairTickerInfo
	KSMEUR    PairTickerInfo
	KSMGBP    PairTickerInfo
	KSMUSD    PairTickerInfo
	KSMXBT    PairTickerInfo
	LINKAUD   PairTickerInfo
	LINKETH   PairTickerInfo
	LINKEUR   PairTickerInfo
	LINKGBP   PairTickerInfo
	LINKUSD   PairTickerInfo
	LINKUSDT  PairTickerInfo
	LINKXBT   PairTickerInfo
	LSKETH    PairTickerInfo
	LSKEUR    PairTickerInfo
	LSKUSD    PairTickerInfo
	LSKXBT    PairTickerInfo
	LTCAUD    PairTickerInfo
	LTCETH    PairTickerInfo
	LTCGBP    PairTickerInfo
	LTCUSDT   PairTickerInfo
	MANAETH   PairTickerInfo
	MANAEUR   PairTickerInfo
	MANAUSD   PairTickerInfo
	MANAXBT   PairTickerInfo
	NANOETH   PairTickerInfo
	NANOEUR   PairTickerInfo
	NANOUSD   PairTickerInfo
	NANOXBT   PairTickerInfo
	OCEANEUR  PairTickerInfo
	OCEANGBP  PairTickerInfo
	OCEANUSD  PairTickerInfo
	OCEANXBT  PairTickerInfo
	OMGETH    PairTickerInfo
	OMGEUR    PairTickerInfo
	OMGUSD    PairTickerInfo
	OMGXBT    PairTickerInfo
	OXTETH    PairTickerInfo
	OXTEUR    PairTickerInfo
	OXTUSD    PairTickerInfo
	OXTXBT    PairTickerInfo
	PAXGETH   PairTickerInfo
	PAXGEUR   PairTickerInfo
	PAXGUSD   PairTickerInfo
	PAXGXBT   PairTickerInfo
	QTUMETH   PairTickerInfo
	QTUMEUR   PairTickerInfo
	QTUMUSD   PairTickerInfo
	QTUMXBT   PairTickerInfo
	REPV2ETH  PairTickerInfo
	REPV2EUR  PairTickerInfo
	REPV2USD  PairTickerInfo
	REPV2XBT  PairTickerInfo
	SCETH     PairTickerInfo
	SCEUR     PairTickerInfo
	SCUSD     PairTickerInfo
	SCXBT     PairTickerInfo
	SNXAUD    PairTickerInfo
	SNXETH    PairTickerInfo
	SNXEUR    PairTickerInfo
	SNXGBP    PairTickerInfo
	SNXUSD    PairTickerInfo
	SNXXBT    PairTickerInfo
	STORJETH  PairTickerInfo
	STORJEUR  PairTickerInfo
	STORJUSD  PairTickerInfo
	STORJXBT  PairTickerInfo
	TBTCETH   PairTickerInfo
	TBTCEUR   PairTickerInfo
	TBTCUSD   PairTickerInfo
	TBTCXBT   PairTickerInfo
	TRXETH    PairTickerInfo
	TRXEUR    PairTickerInfo
	TRXUSD    PairTickerInfo
	TRXXBT    PairTickerInfo
	UNIETH    PairTickerInfo
	UNIEUR    PairTickerInfo
	UNIUSD    PairTickerInfo
	UNIXBT    PairTickerInfo
	USDCAUD   PairTickerInfo
	USDCEUR   PairTickerInfo
	USDCGBP   PairTickerInfo
	USDCHF    PairTickerInfo
	USDCUSD   PairTickerInfo
	USDCUSDT  PairTickerInfo
	USDTAUD   PairTickerInfo
	USDTCAD   PairTickerInfo
	USDTCHF   PairTickerInfo
	USDTEUR   PairTickerInfo
	USDTGBP   PairTickerInfo
	USDTJPY   PairTickerInfo
	USDTZUSD  PairTickerInfo
	WAVESETH  PairTickerInfo
	WAVESEUR  PairTickerInfo
	WAVESUSD  PairTickerInfo
	WAVESXBT  PairTickerInfo
	XBTAUD    PairTickerInfo
	XBTCHF    PairTickerInfo
	XBTDAI    PairTickerInfo
	XBTUSDC   PairTickerInfo
	XBTUSDT   PairTickerInfo
	XDGEUR    PairTickerInfo
	XDGUSD    PairTickerInfo
	XETCXETH  PairTickerInfo
	XETCXXBT  PairTickerInfo
	XETCZEUR  PairTickerInfo
	XETCZUSD  PairTickerInfo
	XETHXXBT  PairTickerInfo
	XETHXXBTD PairTickerInfo
	XETHZCAD  PairTickerInfo
	XETHZCADD PairTickerInfo
	XETHZEUR  PairTickerInfo
	XETHZEURD PairTickerInfo
	XETHZGBP  PairTickerInfo
	XETHZGBPD PairTickerInfo
	XETHZJPY  PairTickerInfo
	XETHZJPYD PairTickerInfo
	XETHZUSD  PairTickerInfo
	XETHZUSDD PairTickerInfo
	XLTCXXBT  PairTickerInfo
	XLTCZEUR  PairTickerInfo
	XLTCZJPY  PairTickerInfo
	XLTCZUSD  PairTickerInfo
	XMLNXETH  PairTickerInfo
	XMLNXXBT  PairTickerInfo
	XMLNZEUR  PairTickerInfo
	XMLNZUSD  PairTickerInfo
	XREPXETH  PairTickerInfo
	XREPXXBT  PairTickerInfo
	XREPZEUR  PairTickerInfo
	XREPZUSD  PairTickerInfo
	XRPAUD    PairTickerInfo
	XRPETH    PairTickerInfo
	XRPGBP    PairTickerInfo
	XRPUSDT   PairTickerInfo
	XTZAUD    PairTickerInfo
	XTZETH    PairTickerInfo
	XTZEUR    PairTickerInfo
	XTZGBP    PairTickerInfo
	XTZUSD    PairTickerInfo
	XTZXBT    PairTickerInfo
	XXBTZCAD  PairTickerInfo
	XXBTZCADD PairTickerInfo
	XXBTZEUR  PairTickerInfo
	XXBTZEURD PairTickerInfo
	XXBTZGBP  PairTickerInfo
	XXBTZGBPD PairTickerInfo
	XXBTZJPY  PairTickerInfo
	XXBTZJPYD PairTickerInfo
	XXBTZUSD  PairTickerInfo
	XXBTZUSDD PairTickerInfo
	XXDGXXBT  PairTickerInfo
	XXLMXXBT  PairTickerInfo
	XXLMZAUD  PairTickerInfo
	XXLMZEUR  PairTickerInfo
	XXLMZGBP  PairTickerInfo
	XXLMZUSD  PairTickerInfo
	XXMRXXBT  PairTickerInfo
	XXMRZEUR  PairTickerInfo
	XXMRZUSD  PairTickerInfo
	XXRPXXBT  PairTickerInfo
	XXRPZCAD  PairTickerInfo
	XXRPZEUR  PairTickerInfo
	XXRPZJPY  PairTickerInfo
	XXRPZUSD  PairTickerInfo
	XZECXXBT  PairTickerInfo
	XZECZEUR  PairTickerInfo
	XZECZUSD  PairTickerInfo
	YFIAUD    PairTickerInfo
	YFIETH    PairTickerInfo
	YFIEUR    PairTickerInfo
	YFIGBP    PairTickerInfo
	YFIUSD    PairTickerInfo
	YFIXBT    PairTickerInfo
	ZEURZUSD  PairTickerInfo
	ZGBPZUSD  PairTickerInfo
	ZUSDZCAD  PairTickerInfo
	ZUSDZJPY  PairTickerInfo
}

// DepositAddressesResponse is the response type of a DepositAddresses query to the Kraken API.
type DepositAddressesResponse []struct {
	Address  string `json:"address"`
	Expiretm string `json:"expiretm"`
	New      bool   `json:"new,omitempty"`
}

// WithdrawResponse is the response type of a Withdraw query to the Kraken API.
type WithdrawResponse struct {
	RefID string `json:"refid"`
}

// WithdrawInfoResponse is the response type showing withdrawal information for a selected withdrawal method.
type WithdrawInfoResponse struct {
	Method string    `json:"method"`
	Limit  big.Float `json:"limit"`
	Amount big.Float `json:"amount"`
	Fee    big.Float `json:"fee"`
}

// GetPairTickerInfo is a helper method that returns given `pair`'s `PairTickerInfo`
func (v *TickerResponse) GetPairTickerInfo(pair string) PairTickerInfo {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(pair)

	return f.Interface().(PairTickerInfo)
}

// PairTickerInfo represents ticker information for a pair
type PairTickerInfo struct {
	// Ask array(<price>, <whole lot volume>, <lot volume>)
	Ask []string `json:"a"`
	// Bid array(<price>, <whole lot volume>, <lot volume>)
	Bid []string `json:"b"`
	// Last trade closed array(<price>, <lot volume>)
	Close []string `json:"c"`
	// Volume array(<today>, <last 24 hours>)
	Volume []string `json:"v"`
	// Volume weighted average price array(<today>, <last 24 hours>)
	VolumeAveragePrice []string `json:"p"`
	// Number of trades array(<today>, <last 24 hours>)
	Trades []int `json:"t"`
	// Low array(<today>, <last 24 hours>)
	Low []string `json:"l"`
	// High array(<today>, <last 24 hours>)
	High []string `json:"h"`
	// Today's opening price
	OpeningPrice float64 `json:"o,string"`
}

// TradesResponse represents a list of the last trades
type TradesResponse struct {
	Last   int64
	Trades []TradeInfo
}

// TradesHistoryResponse represents a list of executed trade
type TradesHistoryResponse struct {
	Trades map[string]TradeHistoryInfo `json:"trades"`
	Count  int                         `json:"count"`
}

// TradeHistoryInfo represents a transaction
type TradeHistoryInfo struct {
	TransactionID string  `json:"ordertxid"`
	PostxID       string  `json:"postxid"`
	AssetPair     string  `json:"pair"`
	Time          float64 `json:"time"`
	Type          string  `json:"type"`
	OrderType     string  `json:"ordertype"`
	Price         float64 `json:"price,string"`
	Cost          float64 `json:"cost,string"`
	Fee           float64 `json:"fee,string"`
	Volume        float64 `json:"vol,string"`
	Margin        float64 `json:"margin,string"`
	Misc          string  `json:"misc"`
}

// TradeInfo represents a trades information
type TradeInfo struct {
	Price         string
	PriceFloat    float64
	Volume        string
	VolumeFloat   float64
	Time          int64
	Buy           bool
	Sell          bool
	Market        bool
	Limit         bool
	Miscellaneous string
}

// LedgersResponse represents an associative array of ledgers infos
type LedgersResponse struct {
	Ledger map[string]LedgerInfo `json:"ledger"`
}

// LedgerInfo Represents the ledger information
type LedgerInfo struct {
	RefID   string    `json:"refid"`
	Time    float64   `json:"time"`
	Type    string    `json:"type"`
	Aclass  string    `json:"aclass"`
	Asset   string    `json:"asset"`
	Amount  big.Float `json:"amount"`
	Fee     big.Float `json:"fee"`
	Balance big.Float `json:"balance"`
}

// OrderTypes for AddOrder
const (
	OTMarket              = "market"
	OTLimit               = "limit"                  // (price = limit price)
	OTStopLoss            = "stop-loss"              // (price = stop loss price)
	OTTakeProfi           = "take-profit"            // (price = take profit price)
	OTStopLossProfit      = "stop-loss-profit"       // (price = stop loss price, price2 = take profit price)
	OTStopLossProfitLimit = "stop-loss-profit-limit" // (price = stop loss price, price2 = take profit price)
	OTStopLossLimit       = "stop-loss-limit"        // (price = stop loss trigger price, price2 = triggered limit price)
	OTTakeProfitLimit     = "take-profit-limit"      // (price = take profit trigger price, price2 = triggered limit price)
	OTTrailingStop        = "trailing-stop"          // (price = trailing stop offset)
	OTTrailingStopLimit   = "trailing-stop-limit"    // (price = trailing stop offset, price2 = triggered limit offset)
	OTStopLossAndLimit    = "stop-loss-and-limit"    // (price = stop loss price, price2 = limit price)
	OTSettlePosition      = "settle-position"
)

// OrderDescription represents an orders description
type OrderDescription struct {
	AssetPair      string `json:"pair"`
	Close          string `json:"close"`
	Leverage       string `json:"leverage"`
	Order          string `json:"order"`
	OrderType      string `json:"ordertype"`
	PrimaryPrice   string `json:"price"`
	SecondaryPrice string `json:"price2"`
	Type           string `json:"type"`
}

// Order represents a single order
type Order struct {
	TransactionID  string           `json:"-"`
	ReferenceID    string           `json:"refid"`
	UserRef        int              `json:"userref"`
	Status         string           `json:"status"`
	OpenTime       float64          `json:"opentm"`
	StartTime      float64          `json:"starttm"`
	ExpireTime     float64          `json:"expiretm"`
	Description    OrderDescription `json:"descr"`
	Volume         string           `json:"vol"`
	VolumeExecuted float64          `json:"vol_exec,string"`
	Cost           float64          `json:"cost,string"`
	Fee            float64          `json:"fee,string"`
	Price          float64          `json:"price,string"`
	StopPrice      float64          `json:"stopprice.string"`
	LimitPrice     float64          `json:"limitprice,string"`
	Misc           string           `json:"misc"`
	OrderFlags     string           `json:"oflags"`
	CloseTime      float64          `json:"closetm"`
	Reason         string           `json:"reason"`
}

// ClosedOrdersResponse represents a list of closed orders, indexed by id
type ClosedOrdersResponse struct {
	Closed map[string]Order `json:"closed"`
	Count  int              `json:"count"`
}

// OrderBookItem is a piece of information about an order.
type OrderBookItem struct {
	Price  float64
	Amount float64
	Ts     int64
}

// UnmarshalJSON takes a json array from kraken and converts it into an OrderBookItem.
func (o *OrderBookItem) UnmarshalJSON(data []byte) error {
	tmpStruct := struct {
		price  string
		amount string
		ts     int64
	}{}
	tmpArray := []interface{}{&tmpStruct.price, &tmpStruct.amount, &tmpStruct.ts}
	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return err
	}

	o.Price, err = strconv.ParseFloat(tmpStruct.price, 64)
	if err != nil {
		return err
	}
	o.Amount, err = strconv.ParseFloat(tmpStruct.amount, 64)
	if err != nil {
		return err
	}
	o.Ts = tmpStruct.ts
	return nil
}

// DepthResponse is a response from kraken to Depth request.
type DepthResponse map[string]OrderBook

// OrderBook contains top asks and bids.
type OrderBook struct {
	Asks []OrderBookItem
	Bids []OrderBookItem
}

// OpenOrdersResponse response when opening an order
type OpenOrdersResponse struct {
	Open  map[string]Order `json:"open"`
	Count int              `json:"count"`
}

// AddOrderResponse response when adding an order
type AddOrderResponse struct {
	Description    OrderDescription `json:"descr"`
	TransactionIds []string         `json:"txid"`
}

// CancelOrderResponse response when cancelling and order
type CancelOrderResponse struct {
	Count   int  `json:"count"`
	Pending bool `json:"pending"`
}

// QueryOrdersResponse response when checking all orders
type QueryOrdersResponse map[string]Order

// NewOHLC constructor for OHLC
func NewOHLC(input []interface{}) (*OHLC, error) {
	if len(input) != 8 {
		return nil, fmt.Errorf("the length is not 8 but %d", len(input))
	}

	tmp := new(OHLC)
	tmp.Time = time.Unix(int64(input[0].(float64)), 0)
	tmp.Open, _ = strconv.ParseFloat(input[1].(string), 64)
	tmp.High, _ = strconv.ParseFloat(input[2].(string), 64)
	tmp.Low, _ = strconv.ParseFloat(input[3].(string), 64)
	tmp.Close, _ = strconv.ParseFloat(input[4].(string), 64)
	tmp.Vwap, _ = strconv.ParseFloat(input[5].(string), 64)
	tmp.Volume, _ = strconv.ParseFloat(input[6].(string), 64)
	tmp.Count = int(input[7].(float64))

	return tmp, nil
}

// OHLC represents the "Open-high-low-close chart"
type OHLC struct {
	Time   time.Time `json:"time"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
	Vwap   float64   `json:"vwap"`
	Volume float64   `json:"volume"`
	Count  int       `json:"count"`
}

// OHLCResponse represents the OHLC's response
type OHLCResponse struct {
	Pair string  `json:"pair"`
	OHLC []*OHLC `json:"OHLC"`
	Last float64 `json:"last"`
}
