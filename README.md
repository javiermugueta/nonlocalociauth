# 

## export this variables to test the go code in local
```
export tenancy="ocid1.tenancy.oc1..aaaa...ua"
export user="ocid1.user.oc1..aaaaa...wa"
export region="eu-frankfurt-1"
export fingerprint="85:...:d6"
export ppkfile=$(cat ./myppk)
export ppkpasswd="-"

go run nonlocalociauth.go 
```
## container image build ans push
```
docker build -t yourrepo/nonlocalociauth .
docker push yourrepo/nonlocalociauth 
```

## create configmap for k8s
```
kubectl create configmap myppk --from-file=myppk=./myppk 
kubectl get configmaps myppk -o yaml
```

## deploy in k8s
```
kubectl delete deployment nonlocalociauth
kubectl apply -f nonlocalociauth.yaml
```

## helper
```
git add . &&  git commit -m . && git push
```
