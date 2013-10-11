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
define("esri/layers/GridLayout",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/has","esri/kernel","esri/SpatialReference","esri/geometry/Extent","esri/geometry/Polyline"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9=_1(null,{declaredClass:"esri.layers._GridLayout",constructor:function(_a,_b,_c,_d){this.origin=_a;this.cellWidth=_b.width;this.cellHeight=_b.height;this.mapWidth=_c.width;this.mapHeight=_c.height;this.srInfo=_d;},setResolution:function(_e){this._resolution=(_e.xmax-_e.xmin)/this.mapWidth;if(this.srInfo){var _f=Math.round((2*this.srInfo.valid[1])/this._resolution),_10=Math.round(_f/this.cellWidth);this._frameStats=[_10,0,_10-1];}},getCellCoordinates:function(_11){var res=this._resolution,_12=this.origin;return {row:Math.floor((_12.y-_11.y)/(this.cellHeight*res)),col:Math.floor((_11.x-_12.x)/(this.cellWidth*res))};},normalize:function(col){var _13=this._frameStats;if(_13){var _14=_13[0],_15=_13[1],_16=_13[2];if(col<_15){col=col%_14;col=col<_15?col+_14:col;}else{if(col>_16){col=col%_14;}}}return col;},intersects:function(_17,_18){var _19=this.srInfo;if(_19){return _3.some(_18._getParts(_19),function(_1a){return _17.intersects(_1a.extent);});}else{return _17.intersects(_18);}},getCellExtent:function(row,col){var res=this._resolution,_1b=this.origin,_1c=this.cellWidth,_1d=this.cellHeight;return new _7((col*_1c*res)+_1b.x,_1b.y-((row+1)*_1d*res),((col+1)*_1c*res)+_1b.x,_1b.y-(row*_1d*res),new _6(_1b.spatialReference.toJson()));},getLatticeID:function(_1e){var _1f=this.getCellCoordinates({x:_1e.xmin,y:_1e.ymax}),_20=this.getCellCoordinates({x:_1e.xmax,y:_1e.ymin}),_21=_1f.row,_22=_20.row,_23=this.normalize(_1f.col),_24=this.normalize(_20.col);return _21+"_"+_22+"_"+_23+"_"+_24;},sorter:function(a,b){return (a<b)?-1:1;},getCellsInExtent:function(_25,_26){var _27=this.getCellCoordinates({x:_25.xmin,y:_25.ymax}),_28=this.getCellCoordinates({x:_25.xmax,y:_25.ymin}),_29=_27.row,_2a=_28.row,_2b=_27.col,_2c=_28.col,_2d=[],i,j,nj,_2e=[],_2f=[],len,_30,_31,_32,_33,_34=[],_35,_36;for(i=_29;i<=_2a;i++){for(j=_2b;j<=_2c;j++){nj=this.normalize(j);_25=this.getCellExtent(i,nj);_2d.push({row:i,col:nj,extent:_25,resolution:this._resolution});if(_26){_2e.push(_25.xmin,_25.xmax);_2f.push(_25.ymin,_25.ymax);}}}_2b=this.normalize(_2b);_2c=this.normalize(_2c);_2e.sort(this.sorter);_2f.sort(this.sorter);len=_2e.length;for(i=len-1;i>=0;i--){if(i<(len-1)){if(_2e[i]===_2e[i+1]){_2e.splice(i,1);}}}len=_2f.length;for(i=len-1;i>=0;i--){if(i<(len-1)){if(_2f[i]===_2f[i+1]){_2f.splice(i,1);}}}if(_2e.length&&_2f.length){_30=_2e[0];_31=_2e[_2e.length-1];_32=_2f[0];_33=_2f[_2f.length-1];len=_2e.length;for(i=0;i<len;i++){_34.push([[_2e[i],_33],[_2e[i],_32]]);}len=_2f.length;for(i=0;i<len;i++){_34.push([[_30,_2f[i]],[_31,_2f[i]]]);}_35=new _8({paths:_34,spatialReference:this.origin.spatialReference.toJson()});_36=_29+"_"+_2a+"_"+_2b+"_"+_2c;_2d.push({latticeID:_36,lattice:_35,resolution:this._resolution});}return {minRow:_29,maxRow:_2a,minCol:_2b,maxCol:_2c,cells:_2d};}});if(_4("extend-esri")){_2.setObject("layers._GridLayout",_9,_5);}return _9;});