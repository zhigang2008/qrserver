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
define("esri/layers/Field",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/layers/RangeDomain","esri/layers/CodedValueDomain"],function(_1,_2,_3,_4,_5,_6){var _7=_1(null,{declaredClass:"esri.layers.Field",constructor:function(_8){if(_8&&_2.isObject(_8)){this.name=_8.name;this.type=_8.type;this.alias=_8.alias;this.length=_8.length;this.editable=_8.editable;this.nullable=_8.nullable;var _9=_8.domain;if(_9&&_2.isObject(_9)){switch(_9.type){case "range":this.domain=new _5(_9);break;case "codedValue":this.domain=new _6(_9);break;}}}}});if(_3("extend-esri")){_2.setObject("layers.Field",_7,_4);}return _7;});