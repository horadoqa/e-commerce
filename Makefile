# =====================================
# Docker Makefile (Ubuntu / WSL)
# =====================================

SHELL := /bin/bash

INSTALL_SCRIPT := ./install-docker-ubuntu.sh

.PHONY: help install status start stop restart

help:
	@echo ""
	@echo "-----------------------------------"
	@echo "             DOCKER                "
	@echo "-----------------------------------"
	@echo "  make install   -> Instala o Docker no WSL2"
	@echo "  make status    -> Exibe o status do serviço"
	@echo "  make start     -> Inicia o serviço Docker"
	@echo "  make stop      -> Para o serviço Docker"
	@echo "  make restart   -> Reinicia o serviço Docker"
	@echo "  make version   -> Verifica a versão do Docker"
	@echo ""
	@echo "-----------------------------------"
	@echo "        Banco de Dados             "
	@echo "-----------------------------------"
	@echo "  make create-db   -> Cria o banco de Dados"
	@echo "  make remove-db   -> Apaga o banco de Dados"

install:
	@echo "==> Instalando Docker..."
	sudo bash $(INSTALL_SCRIPT)

status:
	@echo "==> Status do Docker..."
	sudo systemctl status docker

start:
	@echo "==> Iniciando Docker..."
	sudo systemctl start docker

stop:
	@echo "==> Parando Docker..."
	sudo systemctl stop docker

restart:
	@echo "==> Reiniciando Docker..."
	sudo systemctl restart docker

version:
	@echo "==> Verificando a versão do Docker..."
	docker --version

create-db:
	@echo "==> Cria o Banco de Dados ecommerce-postgres..."
	cd DB && docker compose up -d

remover-db:
	@echo "==> Remover o Banco de Dados ecommerce-postgres..."
	cd DB && docker compose down
