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
define("esri/dijit/editing/tools/ToggleToolBase",["dojo/_base/declare","dojo/_base/lang","dojo/has","dijit/form/ToggleButton","esri/dijit/editing/tools/ToolBase","esri/kernel"],function(_1,_2,_3,_4,_5,_6){var _7=_1([_4,_5],{declaredClass:"esri.dijit.editing.tools.ToggleToolBase",postCreate:function(){this.inherited(arguments);if(this._setShowLabelAttr){this._setShowLabelAttr(false);}},destroy:function(){_4.prototype.destroy.apply(this,arguments);_5.prototype.destroy.apply(this,arguments);},setChecked:function(_8){_4.prototype.setChecked.apply(this,arguments);}});if(_3("extend-esri")){_2.setObject("dijit.editing.tools.ToggleToolBase",_7,_6);}return _7;});