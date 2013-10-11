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
define("esri/dijit/editing/toolbars/ToolbarBase",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/connect","dojo/has","dijit/Toolbar","dijit/ToolbarSeparator","esri/kernel","esri/lang"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9){var _a=_1([_6],{declaredClass:"esri.dijit.editing.toolbars.ToolbarBase",_enabled:true,graphicsAdded:function(){},drawEnd:function(){},onApplyEdits:function(){},onDelete:function(){},constructor:function(_b,_c){if(!_b||!_b.settings){return;}this._tools=[];this._tbConnects=[];this._initialize(_b.settings);},postCreate:function(){this._createTools();this.deactivate();},destroy:function(){var _d;var _e=this._tools;for(_d in _e){if(_e.hasOwnProperty(_d)&&_9.isDefined(this._tools[_d])){this._tools[_d].destroy();}}_3.forEach(this._tbConnects,_4.disconnect);this.inherited(arguments);},activate:function(_f){this._enabled=true;},deactivate:function(){var _10;this._enabled=false;this._layer=null;this._geometryType=null;var _11=this._tools;for(_10 in _11){if(_11.hasOwnProperty(_10)){this._tools[_10].deactivate();this._tools[_10].setChecked(false);}}},isEnabled:function(){return this._enabled;},setActiveSymbol:function(_12){this._activeSymbol=_12;},_getSymbol:function(){},_createTools:function(){},_initialize:function(_13){this._settings=_13;this._toolbar=_13.drawToolbar;this._editToolbar=_13.editToolbar;this._initializeToolbar();},_activateTool:function(_14,_15){if(this._activeTool){this._activeTool.deactivate();}if(_15===true&&this._activeTool==this._tools[_14]){this._activeTool.setChecked(false);this._activeTool=null;}else{this._activeTool=this._tools[_14];this._activeTool.setChecked(true);this._activeTool.activate(null);}},_createSeparator:function(){this.addChild(new _7());}});if(_5("extend-esri")){_2.setObject("dijit.editing.toolbars.ToolbarBase",_a,_8);}return _a;});