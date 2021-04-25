export tenancy="ocid1.tenancy.oc1..aaaa...ua"
export user="ocid1.user.oc1..aaaaa...wa"
export region="eu-frankfurt-1"
export fingerprint="85:...:d6"
export ppkfile=$(cat ./myppk)
export ppkpasswd="-"

go run nonlocalociauth.go 

docker build -t javiermugueta/nonlocalociauth .
docker push javiermugueta/nonlocalociauth 

kubectl create configmap myppk --from-file=myppk=./myppk 
kubectl get configmaps myppk -o yaml

kubectl delete deployment nonlocalociauth
kubectl apply -f nonlocalociauth.yaml



