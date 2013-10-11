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
define("esri/ServerInfo",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/lang"],function(_1,_2,_3,_4,_5){var _6=_1(null,{declaredClass:"esri.ServerInfo",constructor:function(_7){_2.mixin(this,_7);},toJson:function(){return _5.fixJson({server:this.server,tokenServiceUrl:this.tokenServiceUrl,adminTokenServiceUrl:this.adminTokenServiceUrl,shortLivedTokenValidity:this.shortLivedTokenValidity,owningSystemUrl:this.owningSystemUrl,owningTenant:this.owningTenant,currentVersion:this.currentVersion,hasPortal:this.hasPortal,hasServer:this.hasServer});}});if(_3("extend-esri")){_4.ServerInfo=_6;}return _6;});