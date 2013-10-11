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
define("esri/layers/ArcGISTiledMapServiceLayer",["dojo/_base/kernel","dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/json","dojo/has","dojo/io-query","esri/kernel","esri/urlUtils","esri/SpatialReference","esri/layers/TiledMapServiceLayer","esri/layers/ArcGISMapServiceLayer","esri/layers/TileInfo","esri/layers/TimeInfo"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a,_b,_c,_d,_e){var _f=_2([_b,_c],{declaredClass:"esri.layers.ArcGISTiledMapServiceLayer",_agolAttrs:["World_Topo_Map","World_Street_Map","Ocean_Basemap"],constructor:function(url,_10){if(_10){if(_10.roundrobin){_1.deprecated(this.declaredClass+" : Constructor option 'roundrobin' deprecated. Use option 'tileServers'.");_10.tileServers=_10.roundrobin;}this._setTileServers(_10.tileServers);this._loadCallback=_10.loadCallback;}this._params=_3.mixin({},this._url.query);this._initLayer=_3.hitch(this,this._initLayer);var _11=_10&&_10.resourceInfo;if(_11){this._initLayer(_11);}else{this._load=_3.hitch(this,this._load);this._load();}this.registerConnectEvents("esri.layers.ArcGISTiledMapServiceLayer",true);},_TILE_FORMATS:{PNG:"png",PNG8:"png",PNG24:"png",PNG32:"png",JPG:"jpg",JPEG:"jpg",GIF:"gif"},_setTileServers:function(_12){if(_12&&_12.length>0){this.tileServers=_12;var i,il=_12.length;for(i=0;i<il;i++){_12[i]=_9.urlToObject(_12[i]).path;}}},_initLayer:function(_13,io){this.inherited(arguments);this.resourceInfo=_5.toJson(_13);this.tileInfo=new _d(_13.tileInfo);if(!this.spatialReference&&this.tileInfo.spatialReference){this.spatialReference=new _a(this.tileInfo.spatialReference.toJson());}this.isPNG32=this.tileInfo.format==="PNG24"||this.tileInfo.format==="PNG32";if(_13.timeInfo){this.timeInfo=new _e(_13.timeInfo);}var _14=this._url.path,_15=this._loadCallback,_16=(window.location.protocol==="file:")?"http:":window.location.protocol,_17=_14.match(/^https?\:\/\/(server|services)\.arcgisonline\.com\/arcgis\/rest\/services\/([^\/]+)\/mapserver/i),_18=_17&&_17[2];if(!this.tileServers){if(_13.tileServers){this._setTileServers(_13.tileServers);}else{var _19=(_14.search(/^https?\:\/\/server\.arcgisonline\.com/i)!==-1),_1a=(_14.search(/^https?\:\/\/services\.arcgisonline\.com/i)!==-1);if(_19||_1a){this._setTileServers([_14,_14.replace((_19?/server\.arcgisonline/i:/services\.arcgisonline/i),(_19?"services.arcgisonline":"server.arcgisonline"))]);}}}if(_18&&_4.indexOf(this._agolAttrs,_18)!==-1){this.hasAttributionData=true;this.attributionDataUrl=this.attributionDataUrl||(_16+"//static.arcgis.com/attribution/"+_18);}this.loaded=true;this.onLoad(this);if(_15){delete this._loadCallback;_15(this);}},getTileUrl:function(_1b,row,col){var ts=this.tileServers,_1c=this._url.query,_1d=(ts?ts[row%ts.length]:this._url.path)+"/tile/"+_1b+"/"+row+"/"+col;if(_1c){_1d+=("?"+_7.objectToQuery(_1c));}var _1e=this._getToken();if(_1e&&(!_1c||!_1c.token)){_1d+=(_1d.indexOf("?")===-1?"?":"&")+"token="+_1e;}return _9.addProxy(_1d);}});if(_6("extend-esri")){_3.setObject("layers.ArcGISTiledMapServiceLayer",_f,_8);}return _f;});