// Minimal bridge mapping Flutter Web to Wails runtime/go
// Exposes window.wails.runtime and window.wails.go compatible methods
(function(){
  if (!window.wails) window.wails = {};
  // Map to existing Wails runtime injected scripts when in dev/prod
  const rt = window.runtime || window.Wails || {};

  function ensure(path, fallback) {
    const parts = path.split('.');
    let cur = window;
    for (const p of parts) {
      cur[p] = cur[p] || {};
      cur = cur[p];
    }
    return Object.assign(cur, fallback || {});
  }

  // runtime
  ensure('wails.runtime', {
    EventsEmit: (...args) => (rt.Events && rt.Events.Emit ? rt.Events.Emit(...args) : undefined),
    EventsOn: (...args) => (rt.Events && rt.Events.On ? rt.Events.On(...args) : undefined),
    WindowSetTitle: (t) => (rt.Window && rt.Window.SetTitle ? rt.Window.SetTitle(t) : undefined),
    LogInfo: (m) => (rt.Log && rt.Log.Info ? rt.Log.Info(m) : console.log(m)),
  });

  // go binding call: window.wails.go.invoke(namespace, method, ...args)
  ensure('wails.go', {
    invoke: async (ns, method, ...args) => {
      const go = window.go || (rt && rt.Go);
      if (go && go[ns] && typeof go[ns][method] === 'function') {
        return go[ns][method](...args);
      }
      // Fallback to global call method if available
      if (rt && typeof rt.Call === 'function') {
        return rt.Call(`${ns}.${method}`, ...args);
      }
      throw new Error('Wails bindings not available');
    }
  });
})();


