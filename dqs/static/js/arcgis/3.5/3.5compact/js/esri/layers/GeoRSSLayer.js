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
define("esri/layers/GeoRSSLayer",["dojo/_base/declare","dojo/_base/lang","dojo/_base/json","dojo/has","esri/kernel","esri/config","esri/request","esri/layers/ServiceGeneratedFeatureCollection"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9=_1([_8],{declaredClass:"esri.layers.GeoRSSLayer",serviceUrl:location.protocol+"//utility.arcgis.com/sharing/rss",constructor:function(_a,_b){if(_6.defaults.geoRSSService){this.serviceUrl=_6.defaults.geoRSSService;}this._createLayer();this.registerConnectEvents("esri.layers.GeoRSSLayer",true);},parse:function(){this._io=_7({url:this.serviceUrl,content:{url:this.url,refresh:this.loaded?true:undefined,outSR:this._outSR?_3.toJson(this._outSR.toJson()):undefined},callbackParamName:"callback"});return this._io;},_initLayer:function(_c){this.inherited(arguments);if(!this.loaded){this.loaded=true;this.onLoad(this);}}});if(_4("extend-esri")){_2.setObject("layers.GeoRSSLayer",_9,_5);}return _9;});