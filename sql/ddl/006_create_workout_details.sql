CREATE TABLE TRACKER.WORKOUT_DETAILS (
    WORKOUT_NAME VARCHAR(50) NOT NULL REFERENCES TRACKER.WORKOUT(NAME) ON DELETE CASCADE,
    GROUP_ID SMALLINT NOT NULL,
    EXERCISE_NAME VARCHAR(50) NOT NULL REFERENCES TRACKER.EXERCISE(NAME) ON DELETE CASCADE,
    SETS SMALLINT NOT NULL,
    REPS SMALLINT NOT NULL,
    WEIGHT SMALLINT,  
    CRET_TS TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    UPDT_TS TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY(WORKOUT_NAME, GROUP_ID, EXERCISE_NAME)
);