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
define("esri/geometry/Point",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/lang","esri/SpatialReference","esri/geometry/Geometry"],function(_1,_2,_3,_4,_5,_6,_7){var _8=6378137,PI=3.141592653589793,_9=57.29577951308232,_a=0.017453292519943;function _b(_c){return _c*_9;};function _d(_e){return _e*_a;};function _f(lng,lat){if(lat>89.99999){lat=89.99999;}else{if(lat<-89.99999){lat=-89.99999;}}var _10=_d(lat);return [_d(lng)*_8,_8/2*Math.log((1+Math.sin(_10))/(1-Math.sin(_10)))];};function _11(x,y,_12){var _13=_b(x/_8);if(_12){return [_13,_b((PI/2)-(2*Math.atan(Math.exp(-1*y/_8))))];}return [_13-(Math.floor((_13+180)/360)*360),_b((PI/2)-(2*Math.atan(Math.exp(-1*y/_8))))];};var _14={type:"point",x:0,y:0};var _15=_1(_7,{declaredClass:"esri.geometry.Point",constructor:function(x,y,_16){_2.mixin(this,_14);if(_2.isArray(x)){this.x=x[0];this.y=x[1];this.spatialReference=y;}else{if(_2.isObject(x)){_2.mixin(this,x);if(_5.isDefined(this.latitude)){this.y=this.latitude;}if(_5.isDefined(this.longitude)){this.x=this.longitude;}if(this.spatialReference){this.spatialReference=new _6(this.spatialReference);}}else{this.x=x;this.y=y;this.spatialReference=_16;}}this.verifySR();},offset:function(x,y){return new this.constructor(this.x+x,this.y+y,this.spatialReference);},setX:function(x){this.x=x;return this;},setY:function(y){this.y=y;return this;},setLongitude:function(_17){var sr=this.spatialReference;if(sr){if(sr._isWebMercator()){this.x=_f(_17,this.y)[0];}else{if(sr.wkid===4326){this.x=_17;}}}return this;},setLatitude:function(_18){var sr=this.spatialReference;if(sr){if(sr._isWebMercator()){this.y=_f(this.x,_18)[1];}else{if(sr.wkid===4326){this.y=_18;}}}return this;},getLongitude:function(){var sr=this.spatialReference,_19;if(sr){if(sr._isWebMercator()){_19=_11(this.x,this.y)[0];}else{if(sr.wkid===4326){_19=this.x;}}}return _19;},getLatitude:function(){var sr=this.spatialReference,_1a;if(sr){if(sr._isWebMercator()){_1a=_11(this.x,this.y)[1];}else{if(sr.wkid===4326){_1a=this.y;}}}return _1a;},update:function(x,y){this.x=x;this.y=y;return this;},normalize:function(){var x=this.x,sr=this.spatialReference;if(sr){var _1b=sr._getInfo();if(_1b){var _1c=_1b.valid[0],_1d=_1b.valid[1],_1e=2*_1d,_1f;if(x>_1d){_1f=Math.ceil(Math.abs(x-_1d)/_1e);x-=(_1f*_1e);}else{if(x<_1c){_1f=Math.ceil(Math.abs(x-_1c)/_1e);x+=(_1f*_1e);}}}}return new this.constructor(x,this.y,sr);},toJson:function(){var _20={x:this.x,y:this.y},sr=this.spatialReference;if(sr){_20.spatialReference=sr.toJson();}return _20;}});_15.lngLatToXY=_f;_15.xyToLngLat=_11;_15.defaultProps=_14;if(_3("extend-esri")){_2.setObject("geometry.Point",_15,_4);_4.geometry.defaultPoint=_14;}return _15;});