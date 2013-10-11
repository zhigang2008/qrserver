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
define("esri/OperationBase",["dojo/_base/declare","dojo/has","esri/kernel"],function(_1,_2,_3){var _4=_1(null,{declaredClass:"esri.OperationBase",type:"not implemented",label:"not implemented",constructor:function(_5){_5=_5||{};if(_5.label){this.label=_5.label;}},performUndo:function(){console.log("performUndo has not been implemented");},performRedo:function(){console.log("performRedo has not been implemented");}});if(_2("extend-esri")){_3.OperationBase=_4;}return _4;});