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
define("esri/layers/FeatureType",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/has","esri/kernel","esri/lang","esri/symbols/jsonUtils","esri/layers/RangeDomain","esri/layers/CodedValueDomain","esri/layers/InheritedDomain","esri/layers/FeatureTemplate"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a,_b){var _c=_1(null,{declaredClass:"esri.layers.FeatureType",constructor:function(_d){if(_d&&_2.isObject(_d)){this.id=_d.id;this.name=_d.name;var _e=_d.symbol;if(_e){this.symbol=_7.fromJson(_e);}var _f=_d.domains,_10,i;var _11=this.domains={};for(_10 in _f){if(_f.hasOwnProperty(_10)){var _12=_f[_10];switch(_12.type){case "range":_11[_10]=new _8(_12);break;case "codedValue":_11[_10]=new _9(_12);break;case "inherited":_11[_10]=new _a(_12);break;}}}var _13=_d.templates;if(_13){var _14=this.templates=[];for(i=0;i<_13.length;i++){_14.push(new _b(_13[i]));}}}},toJson:function(){var _15={id:this.id,name:this.name,symbol:this.symbol&&this.symbol.toJson()};var _16,_17=this.domains,_18=this.templates,_19=_6.fixJson;if(_17){var _1a=_15.domains={};for(_16 in _17){if(_17.hasOwnProperty(_16)){_1a[_16]=_17[_16]&&_17[_16].toJson();}}_19(_1a);}if(_18){_15.templates=_3.map(_18,function(_1b){return _1b.toJson();});}return _19(_15);}});if(_4("extend-esri")){_2.setObject("layers.FeatureType",_c,_5);}return _c;});