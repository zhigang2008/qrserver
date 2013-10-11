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
define("esri/dijit/BasemapLayer",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel"],function(_1,_2,_3,_4){var _5=_1(null,{declaredClass:"esri.dijit.BasemapLayer",constructor:function(_6){_6=_6||{};if(!_6.url&&!_6.type){console.error("esri.dijit.BasemapLayer: unable to find the 'url' or 'type' property in parameters");}this.url=_6.url;this.type=_6.type;this.isReference=(_6.isReference===true)?true:false;this.displayLevels=_6.displayLevels;this.visibleLayers=_6.visibleLayers;this.bandIds=_6.bandIds;this.opacity=_6.opacity;}});if(_3("extend-esri")){_2.setObject("dijit.BasemapLayer",_5,_4);}return _5;});