LogAnalyzer->Process the logs stored in a given log file and stores them and counts the frequency of levels in given file
|
|___>ProcessFile->takes in log filepath and lineperchunk as input
    Divide the logs into chunks and proccess each chunk in seperate go  routine and Aggregate the level count of each
    go routine output using buffered channel and return levelmap
    which consist of level as key and count as value
|
|___>ProcessLogs->Take logs as input and proccess them and pass it
    to the buffered channel
|
|___>AggregateLevel->Aggregates  the levelmap of each routine