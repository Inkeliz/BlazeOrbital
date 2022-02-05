// This takes quite a few shortcuts by hardcoding stuff and not worrying too hard about
// generality and optimization. That's because GPU use isn't the point of this demo.

(() => {

    const vs = `
  attribute vec4 position;
  attribute vec2 satellitePosition;
  attribute vec4 color;
  
  varying vec4 v_color;
  
  void main() {
    gl_PointSize = 5.0;
    gl_Position = (position + vec4(satellitePosition, 0, 0)) * vec4(1.0/10000.0,1.0/10000.0,1,1);
    v_color = color;
  }`;

    const fs = `
  precision mediump float;
  varying vec4 v_color;

  void main() {
    gl_FragColor = v_color;
  }
  `;

    const gl = document.querySelector("canvas").getContext('webgl');
    const ext = gl.getExtension('ANGLE_instanced_arrays');
    if (!ext) {
        throw new error('need ANGLE_instanced_arrays');
    }
    twgl.addExtensionsToContext(gl);

    const programInfo = twgl.createProgramInfo(gl, [vs, fs]);
    const sunRadius = 300;
    const circleVertices = new Float32Array([0, 0].concat(Array.from(Array(33).keys()).flatMap(x => [sunRadius * Math.sin(x * 2 * Math.PI / 32), sunRadius * Math.cos(x * 2 * Math.PI / 32)])));
    const pointVertices = new Float32Array([0, 0]);
    const bufferInfo = twgl.createBufferInfoFromArrays(gl, {
        position: {
            numComponents: 2,
            data: []
        },
        satellitePosition: {
            numComponents: 2,
            data: [],
            divisor: 1,
            stride: 24,
        },
        color: {
            numComponents: 4,
            data: [],
            divisor: 1,
            stride: 24,
            offset: 8,
        }
    });

    twgl.setBuffersAndAttributes(gl, programInfo, bufferInfo);

    twgl.resizeCanvasToDisplaySize(gl.canvas);
    if (gl.canvas.width > gl.canvas.height) {
        const offset = (gl.canvas.width - gl.canvas.height) / 4;
        gl.viewport(-gl.canvas.width, -gl.canvas.width - offset, 2 * gl.canvas.width, 2 * gl.canvas.width - offset);
    } else {
        const offset = (gl.canvas.height - gl.canvas.width) / 4;
        gl.viewport(-gl.canvas.height - offset, -gl.canvas.height, 2 * gl.canvas.height - offset, 2 * gl.canvas.height);
    }

    gl.clearColor(0, 0, 0, 1);

    const stats = window.Stats();
    stats.showPanel(0); // 0: fps, 1: ms, 2: mb, 3+: custom
    stats.dom.style.left = '';
    stats.dom.style.right = '0px';
    stats.dom.style.transform = 'scale(1.5)';
    stats.dom.style.transformOrigin = 'top right';
    stats.dom.style.opacity = '1';
    document.body.appendChild(stats.dom);

    Object.assign(go.importObject.go, {
        "github.com/inkeliz/satellites/rendering.onNextFrame": (sp) => {
            const id = go.mem.getUint32(sp + 8, true);

            window.requestAnimationFrame(() => {
                stats.begin();
                go._values[id]();
                stats.end();
            });
        },
        "github.com/inkeliz/satellites/rendering.renderSuns": (sp) => {
            gl.clear(gl.COLOR_BUFFER_BIT);

            const _slicePointer = go.mem.getUint32(sp + 8, true) + go.mem.getInt32(sp + 8 + 4, true) * 4294967296;
            const _sliceLength = go.mem.getUint32(sp + 8 + 8, true) + go.mem.getInt32(sp + 8 + 8 + 4, true) * 4294967296;

            const data = new Float32Array(go._inst.exports.mem.buffer, _slicePointer, _sliceLength * 6)

            twgl.setAttribInfoBufferFromArray(gl, bufferInfo.attribs.position, circleVertices);
            twgl.setAttribInfoBufferFromArray(gl, bufferInfo.attribs.satellitePosition, data);
            twgl.setAttribInfoBufferFromArray(gl, bufferInfo.attribs.color, data);

            gl.useProgram(programInfo.program);
            ext.drawArraysInstancedANGLE(gl.TRIANGLE_FAN, 0, circleVertices.length, _sliceLength);
        },
        "github.com/inkeliz/satellites/rendering.renderSatellites": (sp) => {
            const _slicePointer = go.mem.getUint32(sp + 8 , true) + go.mem.getInt32(sp + 8 + 4, true) * 4294967296;
            const _sliceLength = go.mem.getUint32(sp + 8 + 8, true) + go.mem.getInt32(sp + 8 + 8 + 4, true) * 4294967296;

            const data = new Float32Array(go._inst.exports.mem.buffer, _slicePointer, _sliceLength * 6)

            twgl.setAttribInfoBufferFromArray(gl, bufferInfo.attribs.position, pointVertices);
            twgl.setAttribInfoBufferFromArray(gl, bufferInfo.attribs.satellitePosition, data);
            twgl.setAttribInfoBufferFromArray(gl, bufferInfo.attribs.color, data);

            gl.useProgram(programInfo.program);
            ext.drawArraysInstancedANGLE(gl.POINTS, 0, 1, _sliceLength);
        },

    })
})();
