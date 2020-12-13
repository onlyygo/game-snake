set CGO_CXXFLAGS=--std=c++11
set CGO_CPPFLAGS=-IC:\Users\Golang\go\clibs\opencv\build\include
set CGO_LDFLAGS=-LC:\Users\Golang\go\clibs\opencv\sources\build\lib -lopencv_core440.dll -lopencv_highgui440.dll -lopencv_imgcodecs440.dll -lopencv_imgproc440.dll -lopencv_calib3d440.dll -lopencv_dnn440.dll -lopencv_features2d440.dll -lopencv_video440.dll -lopencv_objdetect440.dll -lopencv_videoio440.dll
go build -tags customenv %1
