import * as address from "./address";
import * as answer from "./answer";
import * as cache from "./cache";
import * as crypto from "./crypto";
import * as currency from "./currency";
import * as db from "./db";
import * as email from "./email";
import * as emoji from "./emoji";
import * as file from "./file";
import * as forex from "./forex";
import * as geocoding from "./geocoding";
import * as helloworld from "./helloworld";
import * as id from "./id";
import * as image from "./image";
import * as ip from "./ip";
import * as location from "./location";
import * as otp from "./otp";
import * as postcode from "./postcode";
import * as qr from "./qr";
import * as quran from "./quran";
import * as routing from "./routing";
import * as rss from "./rss";
import * as sentiment from "./sentiment";
import * as sms from "./sms";
import * as stock from "./stock";
import * as stream from "./stream";
import * as sunnah from "./sunnah";
import * as thumbnail from "./thumbnail";
import * as time from "./time";
import * as twitter from "./twitter";
import * as url from "./url";
import * as user from "./user";
import * as weather from "./weather";

export class Client {
  constructor(token: string) {
    this.addressService = new address.AddressService(token);
    this.answerService = new answer.AnswerService(token);
    this.cacheService = new cache.CacheService(token);
    this.cryptoService = new crypto.CryptoService(token);
    this.currencyService = new currency.CurrencyService(token);
    this.dbService = new db.DbService(token);
    this.emailService = new email.EmailService(token);
    this.emojiService = new emoji.EmojiService(token);
    this.fileService = new file.FileService(token);
    this.forexService = new forex.ForexService(token);
    this.geocodingService = new geocoding.GeocodingService(token);
    this.helloworldService = new helloworld.HelloworldService(token);
    this.idService = new id.IdService(token);
    this.imageService = new image.ImageService(token);
    this.ipService = new ip.IpService(token);
    this.locationService = new location.LocationService(token);
    this.otpService = new otp.OtpService(token);
    this.postcodeService = new postcode.PostcodeService(token);
    this.qrService = new qr.QrService(token);
    this.quranService = new quran.QuranService(token);
    this.routingService = new routing.RoutingService(token);
    this.rssService = new rss.RssService(token);
    this.sentimentService = new sentiment.SentimentService(token);
    this.smsService = new sms.SmsService(token);
    this.stockService = new stock.StockService(token);
    this.streamService = new stream.StreamService(token);
    this.sunnahService = new sunnah.SunnahService(token);
    this.thumbnailService = new thumbnail.ThumbnailService(token);
    this.timeService = new time.TimeService(token);
    this.twitterService = new twitter.TwitterService(token);
    this.urlService = new url.UrlService(token);
    this.userService = new user.UserService(token);
    this.weatherService = new weather.WeatherService(token);
  }

  addressService: address.AddressService;
  answerService: answer.AnswerService;
  cacheService: cache.CacheService;
  cryptoService: crypto.CryptoService;
  currencyService: currency.CurrencyService;
  dbService: db.DbService;
  emailService: email.EmailService;
  emojiService: emoji.EmojiService;
  fileService: file.FileService;
  forexService: forex.ForexService;
  geocodingService: geocoding.GeocodingService;
  helloworldService: helloworld.HelloworldService;
  idService: id.IdService;
  imageService: image.ImageService;
  ipService: ip.IpService;
  locationService: location.LocationService;
  otpService: otp.OtpService;
  postcodeService: postcode.PostcodeService;
  qrService: qr.QrService;
  quranService: quran.QuranService;
  routingService: routing.RoutingService;
  rssService: rss.RssService;
  sentimentService: sentiment.SentimentService;
  smsService: sms.SmsService;
  stockService: stock.StockService;
  streamService: stream.StreamService;
  sunnahService: sunnah.SunnahService;
  thumbnailService: thumbnail.ThumbnailService;
  timeService: time.TimeService;
  twitterService: twitter.TwitterService;
  urlService: url.UrlService;
  userService: user.UserService;
  weatherService: weather.WeatherService;
}
