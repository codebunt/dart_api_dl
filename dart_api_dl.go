package dart_api_dl
// #include "stdint.h"
// #include "include/dart_api_dl.c"
//
// // Go does not allow calling C function pointers directly. So we are
// // forced to provide a trampoline.
// bool GoDart_PostCObject(Dart_Port_DL port, Dart_CObject* obj) {
//   return Dart_PostCObject_DL(port, obj);
// }
//	typedef struct MPCRound{
//		char *mtype;
//		char *session;
//		char *peerid;
//		char *msg;
//	}MPCRound;
//
//	int64_t GetMPCRound(void **ppMPCRound, char* mtype, char* session, char* peerid,char* msg) {
//		MPCRound *pWork = (MPCRound *)malloc(sizeof(MPCRound));
//		pWork->mtype=mtype;
//		pWork->session=session;
//		pWork->peerid = peerid;
//		pWork->msg = msg;
//
//		*ppMPCRound = pWork;
//
//		int64_t ptr = (int64_t)pWork;
//
//		return ptr;
//	}
//
//	void clearMPCRoundMemory(MPCRound pWork) {
//		free(&pWork.mtype);
//		free(&pWork.session);
//		free(&pWork.peerid);
//		free(&pWork.msg);
//		free(&pWork);
//	}
import "C"
import "unsafe"

func Init(api unsafe.Pointer) {
	if C.Dart_InitializeApiDL(api) != 0 {
		panic("failed to initialize Dart DL C API: version mismatch. " +
			"must update include/ to match Dart SDK version")
	}
}

func SendToPort(port int64, mtype string,sessionid string, peerid string , msg int64) {
	var obj C.Dart_CObject
	obj._type = C.Dart_CObject_kInt64
	var pwork unsafe.Pointer
	ptrAddr := C.GetMPCRound(&pwork, C.CString(mtype), C.CString(sessionid),C.CString(peerid),C.CString(msg))
	*(*C.int64_t)(unsafe.Pointer(&obj.value[0])) = ptrAddr
	C.GoDart_PostCObject(C.int64_t(port), &obj)
}

func FreeMPCRoundMemory(pointer *int64) {
	ptr := (*C.struct_MPCRound)(unsafe.Pointer(pointer))
	C.clearMPCRoundMemory(*ptr);
}