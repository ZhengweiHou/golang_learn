
``` mermaid
sequenceDiagram
    Client.Do->>Client.do:1
```

``` go
Client.Do
Client.do
    ...
    // 拿到 deadline，没指定时，默认是 0
    deadline = c.deadline()
    ...
    -> c.send(req, deadline)

=> c.send(req, deadline)
    ...
    // 拿到 RoundTripper实现，没指定时，默认是 http.DefaultTransport
    rt = -> c.transport() 
    -> send(req, rt, deadline)
    ...

=> send(req, rt, deadline)
    ...
    // 执行 RoundTripper(默认是http.DefaultTransport) 的 RoundTrip 方法
    -> rt.RoundTrip(req)  // 暂不讨论自定义的 RoundTripper
        -> Transport.RoundTrip(req)
            -> Transport.roundTrip()
    ...

=> Transport.RoundTrip(req) -> Transport.roundTrip(req)
    -> once.Do: t.onceSetNextProtoDefaults()
    // 获取trace，用于汇报trace TODO trace怎么指定实现默认实现是空
    -> trace := httptrace.ContextClientTrace(ctx)
    ...
    if isHTTP
        validateHeaders
    origReq := req
	req = -> setupRewindBody(req) // ?? 重新构建了req??
    validMethod(req.Method)
    ...ctx.cancel相关的处理
    for{
        select {
		case <-ctx.Done():
			req.closeBody()
			return nil, context.Cause(ctx)
		default:
		}

        treq := &transportRequest{req, trace, ctx, cancel}
        cm = t.connectMethodForRequest(treq)
        -> pconn = t.getConn(treq, cm) // 建链
        -> resp = pconn.roundTrop(treq)
        resp.Request = origReq
        return resp
    }

=> Transport.getConn(treq, cm)
    trace.GetConn(..)
    w = new wantConn{}
    // 尝试添加到Idel连接获取队列，则添加到等待队列
    delivered = -> t.queueForIdleConn(w) #qc1
    if !delivered
        -> t.queueForDial(w) #qc2
    
    // 等待连接交付
    select {
	case r := <-w.result: // r : connOrError 由 wantConn.tryDeliver() 响应
        trace.GotConn(..)
        return r.pc, r.err
    case <-treq.ctx.Done():
    ｝

#qc1=> Transport.queueForIdleConn(w)
    if t.DisableKeepAlives // 不允许keepalive
        return false

	t.idleMu.Lock() defer t.idleMu.Unlock() // idle锁
    t.closeIdle = false // TODO 为什么？难道因为
    // 计算过期时间
    oldTime = now - t.IdleConnTimeout
    if list, ok := t.idleConn[w.key]; ok { // 存在w.key对应的空闲连接列表
        pconn := range list{
            if 太旧
                go pconn.closeConnIfStillIdle()
            if 损坏||太旧
                idle list移除连接
                continue
            // 尝试交付当前闲置连接
			-> delivered = w.tryDeliver(pconn, nil, pconn.idleAt)
            if delivered {
				if pconn.alt != nil {
					// HTTP/2: multiple clients can share pconn.
					// Leave it in the list.
				} else {
					// HTTP/1: only one client can use pconn.
					// Remove it from the list.
					t.idleLRU.remove(pconn)
					list = list[:len(list)-1]
				}
			}
			stop = true
            return delivered
        } // range list
    } // 存在空闲连接列表

    q := t.idleConnWait[w.key]
    q.pushBack(w)
    return false // delivered = false


    


#qc2=> Transport.queueForDial(w)
    w.beforeDial() // TODO 怎么指定实现，作用是什么？
    if t.MaxConnsPerHost <= 0  // 没有限制
        t.startDialConnForLocked(w)
		return
    if 池子(t.connsPerHost[w.key])足够
        t.startDialConnForLocked(w)
        池子计数+1
        return
    q := t.connsPerHostWait[w.key] // 等待队列
    q.push(w) // 添加到等待队列 TODO 等待队列从哪里处理

=> Transport.startDialConnForLocked(w)
    t.dialsInProgress.cleanFrontCanceled()
	t.dialsInProgress.pushBack(w)    
    go t.dialConnFor(w)

=> Transport.dialConnFor(w)
    defer w.afterDial()
    ...
    *-> pc, err := t.dialConn(ctx, w.cm) // 建链
    ...
```
