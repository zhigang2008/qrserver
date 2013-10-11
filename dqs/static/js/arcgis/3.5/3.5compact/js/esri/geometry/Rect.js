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
define("esri/geometry/Rect",["dojo/_base/declare","dojo/_base/lang","dojo/has","dojox/gfx/_base","esri/kernel","esri/SpatialReference","esri/geometry/Geometry","esri/geometry/Point","esri/geometry/Extent"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9){function _a(_b){return new _9(parseFloat(_b.x),parseFloat(_b.y)-parseFloat(_b.height),parseFloat(_b.x)+parseFloat(_b.width),parseFloat(_b.y),_b.spatialReference);};var _c=_1(_7,{declaredClass:"esri.geometry.Rect",constructor:function(_d,y,_e,_f,_10){_2.mixin(this,_4.defaultRect);if(_2.isObject(_d)&&_d.type==="extent"){y=_d.ymax;_e=_d.getWidth();_f=_d.getHeight();_10=_d.spatialReference;_d=_d.xmin;}if(_2.isObject(_d)){_2.mixin(this,_d);if(this.spatialReference){this.spatialReference=new _6(this.spatialReference);}}else{this.x=_d;this.y=y;this.width=_e;this.height=_f;this.spatialReference=_10;}this.verifySR();},getCenter:function(){return new _8(this.x+this.width/2,this.y+this.height/2,this.spatialReference);},offset:function(ox,oy){return new _c(this.x+ox,this.y+oy,this.width,this.height,this.spatialReference);},intersects:function(_11){if((_11.x+_11.width)<=this.x){return false;}if((_11.y+_11.height)<=this.y){return false;}if(_11.y>=(this.y+this.height)){return false;}if(_11.x>=(this.x+this.width)){return false;}return true;},getExtent:function(){return _a(this);},update:function(x,y,_12,_13,_14){this.x=x;this.y=y;this.width=_12;this.height=_13;this.spatialReference=_14;return this;}});if(_3("extend-esri")){_2.setObject("geometry.Rect",_c,_5);_5.geometry._rectToExtent=_a;_5.geometry._extentToRect=function(_15){return new _c(_15);};}return _c;});