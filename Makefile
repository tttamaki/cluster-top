BIN_DIR := /usr/local/bin
SYSTEMD_DIR := /etc/systemd/system

all:
	cd proc; go install
	go build cluster-top-node.go config.go data.go
	go build cluster-top-router.go config.go data.go
	go build cluster-top.go config.go data.go
	go build cluster-top-local.go config.go data.go
clean:
	cd proc; go clean
	go clean
install:
	install -v cluster-top-local cluster-top-node cluster-top-router cluster-top $(BIN_DIR)
	install docs/cluster-top-node.service $(SYSTEMD_DIR)
	systemctl enable cluster-top-node.service
	systemctl start cluster-top-node.service
uninstall:
	systemctl disable cluster-top-node.service
	systemctl stop cluster-top-node.service
	rm -f $(SYSTEMD_DIR)/cluster-top-node.service 
	rm -f $(BIN_DIR)/cluster-top-local
	rm -f $(BIN_DIR)/cluster-top-node
	rm -f $(BIN_DIR)/cluster-top-router
	rm -f $(BIN_DIR)/cluster-top


install-router:
	install docs/cluster-top-router.service $(SYSTEMD_DIR)
	systemctl enable cluster-top-router.service
	systemctl start cluster-top-router.service

uninstall-router:
	systemctl disable cluster-top-router.service
	systemctl stop cluster-top-router.service
	rm -f $(SYSTEMD_DIR)/cluster-top-router.service 
