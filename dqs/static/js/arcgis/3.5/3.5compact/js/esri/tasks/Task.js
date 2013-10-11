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
define("esri/tasks/Task",["dojo/_base/declare","dojo/_base/lang","dojo/_base/json","dojo/has","esri/kernel","esri/deferredUtils","esri/urlUtils","esri/Evented"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9=_1(_8,{declaredClass:"esri.tasks._Task",constructor:function(_a){if(_a&&_2.isString(_a)){this._url=_7.urlToObject(this.url=_a);}this.normalization=true;this._errorHandler=_2.hitch(this,this._errorHandler);this.registerConnectEvents("esri.tasks._Task",{"error":["error"]});},_useSSL:function(){var _b=this._url,re=/^http:/i,_c="https:";if(this.url){this.url=this.url.replace(re,_c);}if(_b&&_b.path){_b.path=_b.path.replace(re,_c);}},_encode:function(_d,_e,_f){var _10,_11,_12={},i,p,pl;for(i in _d){if(i==="declaredClass"){continue;}_10=_d[i];_11=typeof _10;if(_10!==null&&_10!==undefined&&_11!=="function"){if(_2.isArray(_10)){_12[i]=[];pl=_10.length;for(p=0;p<pl;p++){_12[i][p]=this._encode(_10[p]);}}else{if(_11==="object"){if(_10.toJson){var _13=_10.toJson(_f&&_f[i]);if(_10.declaredClass==="esri.tasks.FeatureSet"){if(_13.spatialReference){_13.sr=_13.spatialReference;delete _13.spatialReference;}}_12[i]=_e?_13:_3.toJson(_13);}}else{_12[i]=_10;}}}}return _12;},_successHandler:function(_14,_15,_16,dfd){if(_15){this[_15].apply(this,_14);}if(_16){_16.apply(null,_14);}if(dfd){_6._resDfd(dfd,_14);}},_errorHandler:function(err,_17,dfd){this.onError(err);if(_17){_17(err);}if(dfd){dfd.errback(err);}},setNormalization:function(_18){this.normalization=_18;},onError:function(){}});if(_4("extend-esri")){_5.Task=_9;}return _9;});