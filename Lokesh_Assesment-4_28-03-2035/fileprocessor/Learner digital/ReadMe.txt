ConcurrentFileProcessor
|
|____>TextFileExtractor  ->Extract the text files the folder
|
|____>Create Worker->Create go routines for processing each files
|
|____>Start->Starts and manages all tasks for processing
|
|____>Aggregator->Calls the appropriate aggregator based on mode
|
|____>WordCount Aggregator,LineFilterAggregator,ApiCallAggregato
      ->it  process calls their specific utils


utils
|____>ReadDir->Read the Directory and return list of os.Directory
|
|____>ReadFile->read the file and return string
|
|____>WordCount->Counts the words in the LineFilterAggregator
|
|____>FilterLine->Check whether line contains the given word or not and return bool
|
|____>ApiCaller->Post the line to given api and if it fails  it retry for given try count

