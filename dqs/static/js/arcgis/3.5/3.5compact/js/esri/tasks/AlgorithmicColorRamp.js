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
define("esri/tasks/AlgorithmicColorRamp",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/symbols/Symbol","esri/tasks/ColorRamp"],function(_1,_2,_3,_4,_5,_6){var _7=_1(_6,{declaredClass:"esri.tasks.AlgorithmicColorRamp",type:"algorithmic",fromColor:null,toColor:null,algorithm:null,toJson:function(){var _8;switch(this.algorithm.toLowerCase()){case "cie-lab":_8="esriCIELabAlgorithm";break;case "hsv":_8="esriHSVAlgorithm";break;case "lab-lch":_8="esriLabLChAlgorithm";break;default:}var _9={type:"algorithmic",algorithm:_8};_9.fromColor=_5.toJsonColor(this.fromColor);_9.toColor=_5.toJsonColor(this.toColor);return _9;}});if(_3("extend-esri")){_2.setObject("tasks.AlgorithmicColorRamp",_7,_4);}return _7;});