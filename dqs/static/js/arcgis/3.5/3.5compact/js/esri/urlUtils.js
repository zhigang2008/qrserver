/*
 COPYRIGHT 2009 ESRI

 TRADE SECRETS: ESRI PROPRIETARY AND CONFIDENTIAL
 Unpublished material - all rights reserved under the
 Copyright Laws of the United States and applicable international
 laws, treaties, and conventions.

 For additional information, contact:
 Environmental Systems Research Institute, Inc.
 Attn: Contracts and Legal Services Department
 380 New York Street
 Redlands, California, 92373
 USA

 email: contracts@esri.com
 */
//>>built
define("esri/urlUtils",["dojo/_base/lang","dojo/_base/array","dojo/_base/url","dojo/io-query","esri/kernel","esri/config","esri/sniff"],function(_1,_2,_3,_4,_5,_6,_7){var _8={},_9=_6.defaults.io;_8.urlToObject=function(_a){var iq=_a.indexOf("?");if(iq===-1){return {path:_a,query:null};}else{return {path:_a.substring(0,iq),query:_4.queryToObject(_a.substring(iq+1))};}};_8.getProxyUrl=function(_b){var _c=_1.isString(_b)?(_1.trim(_b).toLowerCase().indexOf("https:")===0):_b,_d=_9.proxyUrl,_e,_f,_10,_11,_12="esri.config.defaults.io.proxyUrl is not set.";if(_1.isString(_b)){_11=_8.getProxyRule(_b);if(_11){_d=_11.proxyUrl;}}if(!_d){console.log(_12);throw new Error(_12);}if(_c&&window.location.href.toLowerCase().indexOf("https:")!==0){_f=_d;if(_f.toLowerCase().indexOf("http")!==0){_f=_8.getAbsoluteUrl(_f);}_f=_f.replace(/^http:/i,"https:");if(_8.canUseXhr(_f)){_d=_f;_10=1;}}_e=_8.urlToObject(_d);_e._xo=_10;return _e;};_8.addProxy=function(url){var _13=_8.getProxyRule(url),_14,_15,_16;if(_13){_14=_8.urlToObject(_13.proxyUrl);}else{if(_9.alwaysUseProxy){_14=_8.getProxyUrl();}}if(_14){_15=_8.urlToObject(url);url=_14.path+"?"+_15.path;_16=_4.objectToQuery(_1.mixin(_14.query||{},_15.query));if(_16){url+=("?"+_16);}}return url;};_8.addProxyRule=function(_17){var _18=_17.urlPrefix=(_8.urlToObject(_17.urlPrefix).path).replace(/([^\/])$/,"$1/").replace(/^https?:\/\//ig,"").toLowerCase(),_19=_9.proxyRules,i,len=_19.length,_1a,_1b=len;for(i=0;i<len;i++){_1a=_19[i].urlPrefix;if(_18.indexOf(_1a)===0){if(_18.length===_1a){return -1;}else{_1b=i;break;}}else{if(_1a.indexOf(_18)===0){_1b=i+1;}}}_19.splice(_1b,0,_17);return _1b;};_8.getProxyRule=function(url){var _1c=_9.proxyRules,i,len=_1c.length,_1d=(_8.urlToObject(url).path).replace(/([^\/])$/,"$1/").replace(/^https?:\/\//ig,"").toLowerCase(),_1e;for(i=0;i<len;i++){if(_1d.indexOf(_1c[i].urlPrefix)===0){_1e=_1c[i];break;}}return _1e;};_8.hasSameOrigin=function(_1f,_20,_21){_1f=_1f.toLowerCase();_20=_20.toLowerCase();var _22=window.location.href.toLowerCase();_1f=_1f.indexOf("http")===0?new _3(_1f):(_22=new _3(_22));_20=_20.indexOf("http")===0?new _3(_20):(_1.isString(_22)?new _3(_22):_22);return ((_21||(_1f.scheme===_20.scheme))&&_1f.host===_20.host&&_1f.port===_20.port);};_8.canUseXhr=function(url,_23){var _24=false,_25=_8.hasSameOrigin,_26=_9.corsEnabledServers,_27,_28=-1;if(_7("esri-cors")&&_26&&_26.length){_24=_2.some(_26,function(_29,idx){_27=(_1.trim(_29).toLowerCase().indexOf("http")!==0);if(_25(url,_27?("http://"+_29):_29)||(_27&&_25(url,"https://"+_29))){_28=idx;return true;}return false;});}return _23?_28:_24;};_8.getAbsoluteUrl=function(url){if(_1.isString(url)&&url.indexOf("http://")===-1&&url.indexOf("https://")===-1){if(url.indexOf("//")===0){return window.location.protocol+url;}else{if(url.indexOf("/")===0){return window.location.protocol+"//"+window.location.host+url;}else{return _5._appBaseUrl+url;}}}return url;};if(_7("extend-esri")){_1.mixin(_5,_8);_5._getProxyUrl=_8.getProxyUrl;_5._getProxiedUrl=_8.addProxy;_5._hasSameOrigin=_8.hasSameOrigin;_5._canDoXOXHR=_8.canUseXhr;_5._getAbsoluteUrl=_8.getAbsoluteUrl;}return _8;});