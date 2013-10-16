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
define("esri/layers/ArcGISImageServiceLayer",["dojo/_base/declare","dojo/_base/lang","dojo/_base/Deferred","dojo/_base/array","dojo/_base/json","dojo/_base/config","dojo/has","dojo/io-query","esri/kernel","esri/config","esri/lang","esri/request","esri/deferredUtils","esri/urlUtils","esri/geometry/Extent","esri/layers/MosaicRule","esri/layers/DynamicMapServiceLayer","esri/layers/TimeInfo","esri/layers/Field"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a,_b,_c,_d,_e,_f,_10,_11,_12,_13){var _14=_1(_11,{declaredClass:"esri.layers.ArcGISImageServiceLayer",constructor:function(url,_15){this._url=_e.urlToObject(url);var _16=_15&&_15.imageServiceParameters;this.format=_16&&_16.format;this.interpolation=_16?_16.interpolation:null;this.compressionQuality=_16?_16.compressionQuality:null;this.bandIds=_16?_16.bandIds:null;this.mosaicRule=_16?_16.mosaicRule:null;this.renderingRule=_16?_16.renderingRule:null;this._params=_2.mixin({},this._url.query,{f:"image",interpolation:this.interpolation,format:this.format,compressionQuality:this.compressionQuality,bandIds:this.bandIds?this.bandIds.join(","):null},_16?_16.toJson():{});this._initLayer=_2.hitch(this,this._initLayer);this.useMapImage=(_15&&_15.useMapImage)||false;this._loadCallback=_15&&_15.loadCallback;var _17=_15&&_15.resourceInfo;if(_17){this._initLayer(_17);}else{_c({url:this._url.path,content:_2.mixin({f:"json"},this._url.query),callbackParamName:"callback",load:this._initLayer,error:this._errorHandler});}this.registerConnectEvents("esri.layers.ArcGISImageServiceLayer",{"rendering-change":true});},disableClientCaching:false,_initLayer:function(_18,io){this._findCredential();var ssl=(this.credential&&this.credential.ssl)||(_18&&_18._ssl);if(ssl){this._useSSL();}var _19=this.minScale,_1a=this.maxScale;_2.mixin(this,_18);this.minScale=_19;this.maxScale=_1a;this.initialExtent=(this.fullExtent=this.extent=(new _f(_18.extent)));this.spatialReference=this.initialExtent.spatialReference;this.pixelSizeX=parseFloat(this.pixelSizeX);this.pixelSizeY=parseFloat(this.pixelSizeY);var i,il,_1b=this.minValues,_1c=this.maxValues,_1d=this.meanValues,_1e=this.stdvValues,bs=(this.bands=[]);for(i=0,il=this.bandCount;i<il;i++){bs[i]={min:_1b[i],max:_1c[i],mean:_1d[i],stddev:_1e[i]};}var _1f=this.timeInfo;this.timeInfo=(_1f&&_1f.timeExtent)?new _12(_1f):null;var _20=this.fields=[];var _21=_18.fields;if(_21){for(i=0;i<_21.length;i++){_20.push(new _13(_21[i]));}}this.version=_18.currentVersion;if(!this.version){var ver;if("fields" in _18||"objectIdField" in _18||"timeInfo" in _18){ver=10;}else{ver=9.3;}this.version=ver;}if(_b.isDefined(_18.minScale)&&!this._hasMin){this.setMinScale(_18.minScale);}if(_b.isDefined(_18.maxScale)&&!this._hasMax){this.setMaxScale(_18.maxScale);}var _22={};if(_18.defaultMosaicMethod){_22.method=_18.defaultMosaicMethod;_22.operation=_18.mosaicOperator;_22.sortField=_18.sortField;_22.sortValue=_18.sortValue;}else{_22.method=_10.METHOD_NONE;}this.defaultMosaicRule=new _10(_22);this.defaultMosaicRule.ascending=true;this.loaded=true;this.onLoad(this);var _23=this._loadCallback;if(_23){delete this._loadCallback;_23(this);}},getKeyProperties:function(){var url=this._url.path+"/keyProperties",dfd=new _3(_d._dfdCanceller);if(this.version>10){dfd._pendingDfd=_c({url:url,content:{f:"json"},handleAs:"json",callbackParamName:"callback"});dfd._pendingDfd.then(function(_24){dfd.callback(_24);},function(_25){dfd.errback(_25);});}else{var err=new Error("Layer does not have key properties");err.log=_6.isDebug;dfd.errback(err);}return dfd;},getImageUrl:function(_26,_27,_28,_29){var sr=_26.spatialReference.wkid||_5.toJson(_26.spatialReference.toJson());delete this._params._ts;var _2a=this._url.path+"/exportImage?";_2.mixin(this._params,{bbox:_26.xmin+","+_26.ymin+","+_26.xmax+","+_26.ymax,imageSR:sr,bboxSR:sr,size:_27+","+_28},this.disableClientCaching?{_ts:new Date().getTime()}:{});var _2b=(this._params.token=this._getToken()),_2c=_e.addProxy(_2a+_8.objectToQuery(_2.mixin(this._params,{f:"image"})));if((_2c.length>_a.defaults.io.postLength)||this.useMapImage){this._jsonRequest=_c({url:_2a,content:_2.mixin(this._params,{f:"json"}),callbackParamName:"callback",load:function(_2d,io){var _2e=_2d.href;if(_2b){_2e+=(_2e.indexOf("?")===-1?("?token="+_2b):("&token="+_2b));}_29(_e.addProxy(_2e));},error:this._errorHandler});}else{_29(_2c);}},onRenderingChange:function(){},setInterpolation:function(_2f,_30){this.interpolation=(this._params.interpolation=_2f);if(!_30){this.refresh(true);}},setCompressionQuality:function(_31,_32){this.compressionQuality=(this._params.compressionQuality=_31);if(!_32){this.refresh(true);}},setBandIds:function(ids,_33){var _34=false;if(this.bandIds!==ids){_34=true;}this.bandIds=ids;this._params.bandIds=ids.join(",");if(_34&&!_33){this.onRenderingChange(this._params.bandIds);}if(!_33){this.refresh(true);}},setDefaultBandIds:function(_35){this.bandIds=(this._params.bandIds=null);if(!_35){this.refresh(true);}},setDisableClientCaching:function(_36){this.disableClientCaching=_36;},setMosaicRule:function(_37,_38){this.mosaicRule=_37;this._params.mosaicRule=_5.toJson(_37.toJson());if(!_38){this.refresh(true);}},setRenderingRule:function(_39,_3a){var _3b=false;if(this.renderingRule!==_39){_3b=true;}this.renderingRule=_39;this._params.renderingRule=_5.toJson(_39.toJson());if(_3b&&!_3a){this.onRenderingChange(this._params.renderingRule);}if(!_3a){this.refresh(true);}},setImageFormat:function(_3c,_3d){this.format=(this._params.format=_3c);if(!_3d){this.refresh(true);}},refresh:function(_3e){if(_3e){this.inherited(arguments);}else{var dc=this.disableClientCaching;this.disableClientCaching=true;this.inherited(arguments);this.disableClientCaching=dc;}},exportMapImage:function(_3f,_40){var m=_a.defaults.map,p=_2.mixin({size:m.width+","+m.height},this._params,_3f?_3f.toJson(this.normalization):{},{f:"json"});delete p._ts;this._exportMapImage(this._url.path+"/exportImage",p,_40);}});if(_7("extend-esri")){_2.setObject("layers.ArcGISImageServiceLayer",_14,_9);}return _14;});