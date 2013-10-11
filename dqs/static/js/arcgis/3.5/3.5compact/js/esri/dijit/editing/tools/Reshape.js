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
define("esri/dijit/editing/tools/Reshape",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/connect","dojo/has","esri/layers/FeatureLayer","esri/tasks/query","esri/toolbars/draw","esri/dijit/editing/tools/ToggleToolBase","esri/kernel"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a){var _b=_1([_9],{declaredClass:"esri.dijit.editing.tools.Reshape",id:"btnFeatureReshape",_enabledIcon:"toolbarIcon reshapeIcon",_disabledIcon:"toolbarIcon reshapeIcon",_drawType:_8.POLYLINE,_enabled:true,_label:"NLS_reshapeLbl",activate:function(){_4.disconnect(this._rConnect);this._rConnect=_4.connect(this._toolbar,"onDrawEnd",this,"_onDrawEnd");this.inherited(arguments);},deactivate:function(){this.inherited(arguments);_4.disconnect(this._rConnect);delete this._rConnect;},_onDrawEnd:function(_c){var _d=this._settings.layers;var _e=new _7();_e.geometry=_c;var _f=this._reshapeLayers=_3.filter(_d,function(_10){return (_10.geometryType==="esriGeometryPolygon"||"esriGeometryPolyline");});this._settings.editor._selectionHelper.selectFeatures(_f,_e,_6.SELECTION_NEW,_2.hitch(this,"_reshape",_e));},_reshape:function(_11,_12){var _13=[];var _14=_12;if(_14.length!==1){return;}this._settings.geometryService.reshape(_14[0].geometry,_11.geometry,_2.hitch(this,function(_15){var _16=[_14[0].setGeometry(_15)];this.onApplyEdits([{layer:_14[0].getLayer(),updates:_16}],_2.hitch(this,function(){this._settings.editor._selectionHelper.clearSelection(false);this.onFinished();}));}));}});if(_5("extend-esri")){_2.setObject("dijit.editing.tools.Reshape",_b,_a);}return _b;});