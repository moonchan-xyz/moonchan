type ResolveFunc<V> = (value: V | PromiseLike<V>) => void
type RejectFunc = (reason?: any) => void
type Executor<V> = [ResolveFunc<V>, RejectFunc]

type CacheStatus = 'MISS' | 'DONE' | 'FETCHING' | 'ERROR'
const MISS = 'MISS'
const DONE = 'DONE'
const FETCHING = 'FETCHING'
const ERROR = 'ERROR'


class Soup<K,V> {
  // run this method when data not in cache
  getter: (key:K) => Promise<V>;
  private _cache: Map<K,V> = new Map<K, V>() 
  private _status: Map<K, CacheStatus> = new Map<K, CacheStatus>()
  private _executors = new Map<K, Executor<V>[]>()

  constructor(
    // the raw getter function
    getter : (key:K) => Promise<V>,
    initalMap?: Map<K,V>,
  ){
    this.getter = getter
    initalMap?.forEach((v,k) => {
      this._status.set(k, DONE)
      this._cache.set(k, v)
    })
  }

  // status
  // miss => fetching => done
  // or
  // miss => fetching => error
  private _getStatus(key: K): string {
    const status = this._status.get(key)
    if (typeof status === 'undefined') 
      return MISS
    else 
      return status
  }
  private  _setStatus(key: K, status: CacheStatus) {
    this._status.set(key, status)
  }

  // cache
  private  _hasCache(key: K): boolean {
    return this._cache.has(key)
  }
  private  _getCache(key: K): V | undefined {
    return this._cache.get(key)
  }
  private  _setCache(key: K, value: V) {
    this._cache.set(key, value)
  }

  // executors
  private  _addExecutor(key: K, resolve:ResolveFunc<V>, reject:RejectFunc) {
    if (!this._executors.has(key)) {
      this._executors.set(key, [])
    }
    this._executors.get(key)!.push([resolve, reject])
  }
  private  _delExecutor(key: K) {
    this._executors.delete(key)
  }
  private _getExecutor(key: K): Executor<V>[] {
    const arr = this._executors.get(key)
    if (typeof arr === 'undefined'){
      return []
    }
    this._delExecutor(key)
    return arr
  }

  //
  delAll(key: K) {
    this._status.delete(key)
    this._cache.delete(key)
    this._getExecutor(key).forEach((executor) => {
      executor[1]("deleted")
    })
    this._delExecutor(key)
  }

  // core function
  forceRenew(key: K) {
    this._setStatus(key, FETCHING)
    // this is the original getter
    this.getter(key) // async
    .then((value) => {
      this._setCache(key, value)
      this._setStatus(key, DONE)
      this._getExecutor(key).forEach((executor) => {
        const resolve = executor[0]
        resolve(value)
      })
      this._delExecutor(key)
    })
    .catch((reason) => {
      this._setStatus(key, ERROR)
      console.log(reason)
      this._getExecutor(key).forEach((executor) => {
        const reject = executor[1]
        reject(reason)
      })
      this._delExecutor(key)
    })
  }

  // core function
  renew(key: K) {
    const status = this._getStatus(key)
    if (status === MISS || status === ERROR) {
      this.forceRenew(key)
    } else { // fetching or done
      return
    }
  }

  patch(key: K, value: V) {
    this._setCache(key, value)
    this._setStatus(key, DONE)
  }

  // add call back function
  private _addCallbacks(key: K, resolve: ResolveFunc<V>, reject: RejectFunc) {
    if (this._getStatus(key) === DONE && this._hasCache(key)) {
      // if done, just return what in cache
      resolve(this._getCache(key)!)
    } else {
      // addCache will do nothing when fetching or done
      // if fetching, then add Executor that runs later
      this._addExecutor(key, resolve, reject)
      this.renew(key)
    }
  }

  // interface 
  // usage: someSoup.get(key).then((v)=>{...}).catch(r=>{...})
  get(key: K): Promise<V> {    
    return new Promise((resolve,reject) => {
      this._addCallbacks(key, resolve, reject)
    })
  }

}

export { Soup }