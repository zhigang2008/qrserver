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
define("esri/tasks/AddressCandidate",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/geometry/Point"],function(_1,_2,_3,_4,_5){var _6=_1(null,{declaredClass:"esri.tasks.AddressCandidate",constructor:function(_7){_2.mixin(this,_7);this.location=new _5(this.location);}});if(_3("extend-esri")){_2.setObject("tasks.AddressCandidate",_6,_4);}return _6;});