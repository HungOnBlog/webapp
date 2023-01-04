# Check if Linkerd is installed
if ! command -v linkerd &> /dev/null
then
    echo "Linkerd is not installed"
    echo "Installing Linkerd"
    curl --proto '=https' --tlsv1.2 -sSfL https://run.linkerd.io/install | sh
    echo "Linkerd installed"
else
    echo "Linkerd is already installed"
fi
curl -sL https://linkerd.github.io/linkerd-smi/install | sh

linkerd install --crds | kubectl apply -f -
linkerd install | kubectl apply -f -
linkerd viz install | kubectl apply -f -
linkerd smi install | kubectl apply -f -
