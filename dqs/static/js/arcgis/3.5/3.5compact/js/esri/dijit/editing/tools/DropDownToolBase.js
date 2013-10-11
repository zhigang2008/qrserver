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
define("esri/dijit/editing/tools/DropDownToolBase",["dojo/_base/declare","dojo/_base/lang","dojo/has","dojo/dom-style","dijit/registry","dijit/Menu","dijit/form/ComboButton","esri/dijit/editing/tools/ToolBase","esri/kernel","esri/lang"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a){var _b=_1([_7,_8],{declaredClass:"esri.dijit.editing.tools.DropDownToolBase",_enabled:false,_checked:false,postCreate:function(){this._tools=[];this._createTools();this.inherited(arguments);if(this._setShowLabelAttr){this._setShowLabelAttr(false);}},destroy:function(){var _c;var _d=this._tools;for(_c in _d){if(_d.hasOwnProperty(_c)&&_a.isDefined(_d[_c])){_d[_c].destroy();}}this.inherited(arguments);},_createTools:function(){var i;var _e=new _6();this.dropDown=_e;for(i in this._tools){if(this._tools.hasOwnProperty(i)){_e.addChild(this._tools[i]);}}this._activeTool=_e.getChildren()[0];this._updateUI();},activate:function(_f){this.inherited(arguments);if(!this._activeTool){this._activateDefaultTool();}else{this._activeTool.activate();}},deactivate:function(){this.inherited(arguments);if(this._activeTool){this._activeTool.deactivate();}},enable:function(_10){var _11;for(_11 in this._tools){if(this._tools.hasOwnProperty(_11)){this._tools[_11].enable(_10);}}this.setEnabled(true);this.inherited(arguments);},setChecked:function(_12){this._checked=_12;this._updateUI();},_onDrawEnd:function(_13){},onLayerChange:function(_14,_15,_16){this._activeTool=null;this._activeType=_15;this._activeTemplate=_16;this._activeLayer=_14;},onItemClicked:function(_17,evt){if(this._activeTool){this._activeTool.deactivate();}this._activeTool=_5.byId(_17);if(this._checked===false){this._onClick();}else{this._updateUI();if(this._activeTool){this._activeTool.activate();this._activeTool.setChecked(true);}}},_onClick:function(evt){if(this._enabled===false){return;}this._checked=!this._checked;this.inherited(arguments);},_updateUI:function(){this.attr("disabled",!this._enabled);_4.set(this.focusNode,{outline:"none"});_4.set(this.titleNode,{padding:"0px",border:"none"});if(this._checked){_4.set(this.titleNode,{backgroundColor:"#D4DFF2",border:"1px solid #316AC5"});}else{_4.set(this.titleNode,{backgroundColor:"",border:""});}if(this._activeTool){this.attr("iconClass",this._checked?this._activeTool._enabledIcon:this._activeTool._disabledIcon);this.attr("label",this._activeTool.label);}}});if(_3("extend-esri")){_2.setObject("dijit.editing.tools.DropDownToolBase",_b,_9);}return _b;});