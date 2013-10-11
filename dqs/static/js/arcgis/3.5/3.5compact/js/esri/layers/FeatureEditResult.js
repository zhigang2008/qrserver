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
define("esri/layers/FeatureEditResult",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel"],function(_1,_2,_3,_4){var _5=_1(null,{declaredClass:"esri.layers.FeatureEditResult",constructor:function(_6){if(_6&&_2.isObject(_6)){this.objectId=_6.objectId;this.success=_6.success;if(!_6.success){var _7=_6.error;this.error=new Error();this.error.code=_7.code;this.error.message=_7.description;}}}});if(_3("extend-esri")){_2.setObject("layers.FeatureEditResult",_5,_4);}return _5;});