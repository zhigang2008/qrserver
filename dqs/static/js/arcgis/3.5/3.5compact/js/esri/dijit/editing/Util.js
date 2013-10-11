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
define("esri/dijit/editing/Util",["dojo/_base/lang","dojo/_base/array","dojo/has","esri/kernel"],function(_1,_2,_3,_4){var _5={};_5={findFeatures:function(_6,_7,_8){var _9=_7.objectIdField;var _a=_7.graphics;var _b=_2.filter(_a,function(_c){return _2.some(_6,function(id){return _c.attributes[_9]===id.objectId;});});if(_8){_8(_b);}else{return _b;}},getSelection:function(_d){var _e=[];_2.forEach(_d,function(_f){var _10=_f.getSelectedFeatures();_2.forEach(_10,function(_11){_e.push(_11);});});return _e;}};if(_3("extend-esri")){_1.setObject("dijit.editing.Util.LayerHelper",_5,_4);}return _5;});