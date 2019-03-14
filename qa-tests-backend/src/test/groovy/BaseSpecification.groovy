import com.google.protobuf.Timestamp
import com.jayway.restassured.RestAssured
import common.Constants
import groovy.util.logging.Slf4j
import orchestratormanager.OrchestratorMain
import orchestratormanager.OrchestratorType
import org.junit.Rule
import org.junit.rules.TestName
import org.junit.rules.Timeout
import services.BaseService
import spock.lang.Shared
import spock.lang.Specification
import testrailintegration.TestRailconfig
import util.Env

import java.util.concurrent.TimeUnit

@Slf4j
class BaseSpecification extends Specification {

    @Rule
    Timeout globalTimeout = new Timeout(500000, TimeUnit.MILLISECONDS)
    @Rule
    TestName name = new TestName()
    @Shared
    boolean isTestrail = System.getenv("testrail")
    @Shared
    TestRailconfig tc = new TestRailconfig()
    @Shared
    def resultMap = [:]
    @Shared
    OrchestratorMain orchestrator = OrchestratorType.create(
            Env.mustGetOrchestratorType(),
            Constants.ORCHESTRATOR_NAMESPACE
    )
    @Shared
    private testStartTime

    @Shared
    private dtrId = ""

    def setupSpec() {
        def startTime = System.currentTimeMillis()
        testStartTime = Timestamp.newBuilder().setSeconds(startTime / 1000 as Long)
                .setNanos((int) ((startTime % 1000) * 1000000)).build()
        RestAssured.useRelaxedHTTPSValidation()
        try {
            dtrId = Services.addDockerTrustedRegistry()
            orchestrator.setup()
        } catch (Exception e) {
            println "Error setting up orchestrator: ${e.message}"
            throw e
        }
        /*    if (isTestrail == true) {
            tc.createTestRailInstance()
            tc.setProjectSectionId("Prevent", "Policies")
            tc.createRun()
        }*/
    }
    def setup() { }

    def cleanupSpec() {
        try {
            Services.deleteImageIntegration(dtrId)
            orchestrator.cleanup()
        } catch (Exception e) {
            println "Error to clean up orchestrator: ${e.message}"
            throw e
        }
        /* if (isTestrail == true) {
            List<Integer> caseids = new ArrayList<Integer>(resultMap.keySet());
            tc.updateRun(caseids)
            resultMap.each { entry ->
            println "testcaseId: $entry.key status: $entry.value"
            Integer status = Integer.parseInt(entry.key.toString());
            Integer testcaseId = Integer.parseInt(entry.value.toString());
            tc.addStatusForCase(testcaseId, status);
        }*/
    }

    def cleanup() {
        //Always make sure to revert back to basic auth after each test
        BaseService.useBasicAuth()
    }
}
