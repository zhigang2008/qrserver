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
define("esri/renderers/SymbolAger",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel"],function(_1,_2,_3,_4){var _5=_1(null,{declaredClass:"esri.renderer.SymbolAger",getAgedSymbol:function(_6,_7){},_setSymbolSize:function(_8,_9){switch(_8.type){case "simplemarkersymbol":_8.setSize(_9);break;case "picturemarkersymbol":_8.setWidth(_9);_8.setHeight(_9);break;case "simplelinesymbol":case "cartographiclinesymbol":_8.setWidth(_9);break;case "simplefillsymbol":case "picturefillsymbol":if(_8.outline){_8.outline.setWidth(_9);}break;}}});if(_3("extend-esri")){_2.setObject("renderer.SymbolAger",_5,_4);}return _5;});