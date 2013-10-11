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
define("esri/dijit/editing/tools/Union",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/has","esri/graphicsUtils","esri/graphic","esri/toolbars/draw","esri/dijit/editing/Union","esri/dijit/editing/tools/ButtonToolBase","esri/kernel"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a){var _b=_1([_9],{declaredClass:"esri.dijit.editing.tools.Union",id:"btnFeatureUnion",_enabledIcon:"toolbarIcon unionIcon",_disabledIcon:"toolbarIcon unionIcon",_drawType:_7.POLYLINE,_enabled:true,_label:"NLS_unionLbl",_onClick:function(_c){this._settings.editor._activeTool="UNION";var _d=this._settings.layers;var _e=_3.filter(_d,function(_f){return (_f.geometryType==="esriGeometryPolygon")&&(_f.visible&&_f._isMapAtVisibleScale());});var _10=[];var _11=0;_3.forEach(_e,function(_12,idx){var _13=_12.getSelectedFeatures();if(_13.length>=2){_11++;var _14=_3.map(_13,function(_15){return new _6(_2.clone(_15.toJson()));});this._settings.geometryService.union(_5.getGeometries(_13),_2.hitch(this,function(_16){var _17=[_13.pop().setGeometry(_16)];_10.push({layer:_12,updates:_17,deletes:_13,preUpdates:_14});_11--;if(_11<=0){this.onApplyEdits(_10,_2.hitch(this,function(){if(this._settings.undoRedoManager){var _18=this._settings.undoRedoManager;_3.forEach(this._edits,_2.hitch(this,function(_19){_18.add(new _8({featureLayer:_19.layer,addedGraphics:_19.deletes,preUpdatedGraphics:_19.preUpdates,postUpdatedGraphics:_19.updates}));}),this);}this._settings.editor._selectionHelper.clearSelection(false);this.onFinished();}));}}));}},this);}});if(_4("extend-esri")){_2.setObject("dijit.editing.tools.Union",_b,_a);}return _b;});