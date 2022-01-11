package yamlconfig

var StaticPod = `
apiVersion: v1
kind: Pod
metadata:
  name: ${name}
  labels:
    computedriver: ${name}
spec:
  containers:
    - name: ${name}
      image: ${image}
      imagePullPolicy: IfNotPresent
      resources:
        requests:
          cpu: ${cpu}
          memory: ${memory}
        limits:
          cpu: ${cpu}
          memory: ${memory}
      ports:
        - name: ${name}
          containerPort: ${containerPort}
  hostNetwork: true
`

var StaticService = `
apiVersion: v1
kind: Service
metadata:
  name: ${name}
spec:
  type: NodePort
  ports:
    - port: ${containerPort}
      targetPort: ${containerPort}
      nodePort: ${port}
  selector:
    computedriver: ${name}
`

// func configInit() {
// 	log.Println("static pod temp yaml is {}", StaticPod)
// 	log.Println("static service temp yaml is {}", StaticService)
// }
