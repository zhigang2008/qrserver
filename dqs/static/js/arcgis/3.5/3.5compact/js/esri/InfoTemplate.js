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
define("esri/InfoTemplate",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/lang"],function(_1,_2,_3,_4,_5){var _6=_1(null,{declaredClass:"esri.InfoTemplate",constructor:function(_7,_8){if(_7&&_2.isObject(_7)&&!_2.isFunction(_7)){_2.mixin(this,_7);}else{this.title=_7||"${*}";this.content=_8||"${*}";}},setTitle:function(_9){this.title=_9;return this;},setContent:function(_a){this.content=_a;return this;},toJson:function(){return _5.fixJson({title:this.title,content:this.content});}});if(_3("extend-esri")){_4.InfoTemplate=_6;}return _6;});