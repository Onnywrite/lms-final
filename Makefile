OUTDIR=.\bin
CMD=.\cmd
FRONTEXE=front.exe
ORCHEXE=${ORCH}.exe
AGENTEXE=agent.exe

# shortness
ORCH=orchestrator

.DEFAULT_GOAL := run

build_front:
	go build -o ${OUTDIR}\${FRONTEXE} ${CMD}\main.go

build_orchestrator:
	go build -o ${OUTDIR}\${ORCHEXE} ${CMD}\${ORCH}\${ORCH}.go

build_agent:
	go build -o ${OUTDIR}\${AGENTEXE} ${CMD}\daemon\daemon.go

link_deps:
	xcopy /Y .\resources ${OUTDIR}\resources
	xcopy /Y .\configs ${OUTDIR}\configs

build: build_front build_orchestrator link_deps #build_agent link_deps

run_front: build_front
	${OUTDIR}\${FRONTEXE}

run_orchestrator: build_orchestrator
	${OUTDIR}\${ORCHEXE}

run_agent: build_agent
	${OUTDIR}\${AGENTEXE}

run: build run_orchestrator run_front #run_agent