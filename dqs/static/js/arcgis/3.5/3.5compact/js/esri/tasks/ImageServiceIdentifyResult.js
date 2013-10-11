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
define("esri/tasks/ImageServiceIdentifyResult",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/geometry/jsonUtils","esri/tasks/FeatureSet"],function(_1,_2,_3,_4,_5,_6){var _7=_1(null,{declaredClass:"esri.tasks.ImageServiceIdentifyResult",constructor:function(_8){if(_8.catalogItems){this.catalogItems=new _6(_8.catalogItems);}if(_8.location){this.location=_5.fromJson(_8.location);}this.catalogItemVisibilities=_8.catalogItemVisibilities;this.name=_8.name;this.objectId=_8.objectId;this.value=_8.value;this.properties=_8.properties;}});if(_3("extend-esri")){_2.setObject("tasks.ImageServiceIdentifyResult",_7,_4);}return _7;});