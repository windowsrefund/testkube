## testkube abort

Abort tests or test suites

```
testkube abort <resourceName> [flags]
```

### Options

```
  -h, --help   help for abort
```

### Options inherited from parent commands

```
  -a, --api-uri string     api uri, default value read from config if set (default "https://demo.testkube.io/results")
  -c, --client string      client used for connecting to Testkube API one of proxy|direct (default "proxy")
      --insecure           insecure connection for direct client
      --namespace string   Kubernetes namespace, default value read from config if set (default "testkube")
      --oauth-enabled      enable oauth
      --verbose            show additional debug messages
```

### SEE ALSO

* [testkube](testkube.md)	 - Testkube entrypoint for kubectl plugin
* [testkube abort execution](testkube_abort_execution.md)	 - Aborts execution of the test
* [testkube abort executions](testkube_abort_executions.md)	 - Aborts all executions of the test
* [testkube abort testsuiteexecution](testkube_abort_testsuiteexecution.md)	 - Abort test suite execution
* [testkube abort testsuiteexecutions](testkube_abort_testsuiteexecutions.md)	 - Abort all test suite executions
