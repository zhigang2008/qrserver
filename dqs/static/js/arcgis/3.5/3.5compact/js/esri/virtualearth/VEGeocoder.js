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
define("esri/virtualearth/VEGeocoder",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/Deferred","dojo/has","esri/kernel","esri/urlUtils","esri/tasks/Task","esri/virtualearth/VEGeocodeResult","esri/deferredUtils","esri/request"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a,_b){var _c=_1(_8,{declaredClass:"esri.virtualearth.VEGeocoder",constructor:function(_d){try{_d=_2.mixin({bingMapsKey:null},_d||{});this.url="http://serverapi.arcgisonline.com/veadaptor/production/services/geocode/geocode";this._url=_7.urlToObject(this.url);this._queue=[];this.bingMapsKey=_d.bingMapsKey;this.culture=_d.culture||"en-US";this._errorHandler=_2.hitch(this,this._errorHandler);this._addressToLocationsHandler=_2.hitch(this,this._addressToLocationsHandler);if(!this.bingMapsKey){throw new Error("BingMapsKey must be provided.");}}catch(e){this.onError(e);throw e;}},addressToLocations:function(_e,_f,_10){if(!this.bingMapsKey){console.debug("Server token not retrieved. Queing request to be executed after server token retrieved.");this._queue.push(arguments);return;}var _11=_2.mixin({},this._url.query,{query:_e,token:this.bingMapsKey,culture:this.culture}),_12=this._addressToLocationsHandler,_13=this._errorHandler;var dfd=new _4(_a._dfdCanceller);dfd._pendingDfd=_b({url:this._url.path,content:_11,callbackParamName:"callback",load:function(r,i){_12(r,i,_f,_10,dfd);},error:function(r){_13(r,_10,dfd);}});return dfd;},_addressToLocationsHandler:function(_14,io,_15,_16,dfd){try{_3.forEach(_14,function(_17,i){_14[i]=new _9(_17);});this._successHandler([_14],"onAddressToLocationsComplete",_15,dfd);}catch(err){this._errorHandler(err,_16,dfd);}},onAddressToLocationsComplete:function(){},setBingMapsKey:function(_18){this.bingMapsKey=_18;},setCulture:function(_19){this.culture=_19;}});if(_5("extend-esri")){_2.setObject("virtualearth.VEGeocoder",_c,_6);}return _c;});