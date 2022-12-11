APP=exercise_1_10

SECTION=$(shell pwd)
BIN=$(SECTION)/../../../bin

SOURCE=main.go
EXECUTABLE=$(BIN)/$(APP)
OUTPUT=$(SECTION)/../../../output/$(APP).txt

#Build data file
$(OUTPUT): $(EXECUTABLE)
	$(EXECUTABLE) https://medium.com/swlh/how-quitting-my-corporate-job-for-my-startup-dream-f-cked-my-life-up-3b6b3e29b318 \
                    https://medium.com/swlh/how-quitting-my-corporate-job-for-my-startup-dream-f-cked-my-life-up-3b6b3e29b318 \
					https://medium.com/swlh/how-quitting-my-corporate-job-for-my-startup-dream-f-cked-my-life-up-3b6b3e29b318 \
					https://medium.com/swlh/how-quitting-my-corporate-job-for-my-startup-dream-f-cked-my-life-up-3b6b3e29b318 \
                    https://medium.com/swlh/how-quitting-my-corporate-job-for-my-startup-dream-f-cked-my-life-up-3b6b3e29b318 \
					https://medium.com/swlh/how-quitting-my-corporate-job-for-my-startup-dream-f-cked-my-life-up-3b6b3e29b318 \
					https://medium.com/swlh/how-quitting-my-corporate-job-for-my-startup-dream-f-cked-my-life-up-3b6b3e29b318

# Build app
$(EXECUTABLE): $(SOURCE)
	go build -o $(EXECUTABLE) $(main.go)

run_one:
	$(EXECUTABLE) https://medium.com/swlh/how-quitting-my-corporate-job-for-my-startup-dream-f-cked-my-life-up-3b6b3e29b318

run_many:
	$(EXECUTABLE) https://medium.com/swlh/how-quitting-my-corporate-job-for-my-startup-dream-f-cked-my-life-up-3b6b3e29b318 \
                    https://medium.com/swlh/how-quitting-my-corporate-job-for-my-startup-dream-f-cked-my-life-up-3b6b3e29b318 \
					https://medium.com/swlh/how-quitting-my-corporate-job-for-my-startup-dream-f-cked-my-life-up-3b6b3e29b318 \
					https://medium.com/swlh/how-quitting-my-corporate-job-for-my-startup-dream-f-cked-my-life-up-3b6b3e29b318 \
                    https://medium.com/swlh/how-quitting-my-corporate-job-for-my-startup-dream-f-cked-my-life-up-3b6b3e29b318 \
					https://medium.com/swlh/how-quitting-my-corporate-job-for-my-startup-dream-f-cked-my-life-up-3b6b3e29b318 \
					https://medium.com/swlh/how-quitting-my-corporate-job-for-my-startup-dream-f-cked-my-life-up-3b6b3e29b318

